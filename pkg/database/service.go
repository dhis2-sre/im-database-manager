package database

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"

	"github.com/dhis2-sre/im-database-manager/pkg/model"

	"github.com/dhis2-sre/im-database-manager/internal/apperror"
	"github.com/dhis2-sre/im-database-manager/pkg/config"
	"github.com/dhis2-sre/im-database-manager/pkg/storage"
	jobClient "github.com/dhis2-sre/im-job/pkg/client"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
)

type Service interface {
	Create(d *model.Database) error
	Copy(id uint, d *model.Database, group *models.Group) error
	FindById(id uint) (*model.Database, error)
	Lock(id uint, instanceId uint, userId uint) (*model.Lock, error)
	Unlock(id uint) error
	Upload(d *model.Database, group *models.Group, file io.Reader) (*model.Database, error)
	Download(id uint, dst io.Writer, headers func(contentLength int64)) error
	Delete(id uint) error
	List(groups []*models.Group) ([]*model.Database, error)
	Update(d *model.Database) error
	CreateExternalDownload(databaseID uint, expiration time.Time) (model.ExternalDownload, error)
	FindExternalDownload(uuid uuid.UUID) (model.ExternalDownload, error)
}

type service struct {
	c          config.Config
	s3Client   storage.S3Client
	jobClient  jobClient.Client
	repository Repository
}

func NewService(c config.Config, s3Client storage.S3Client, jobClient jobClient.Client, repository Repository) *service {
	return &service{c, s3Client, jobClient, repository}
}

func (s service) Create(d *model.Database) error {
	return s.repository.Create(d)
}

func (s service) Copy(id uint, d *model.Database, group *models.Group) error {
	source, err := s.FindById(id)
	if err != nil {
		if err.Error() == "record not found" {
			idStr := strconv.FormatUint(uint64(id), 10)
			err = apperror.NewNotFound("database not found", idStr)
		}
		return err
	}

	u, err := url.Parse(source.Url)
	if err != nil {
		return err
	}

	sourceKey := strings.TrimPrefix(u.Path, "/")
	destinationKey := fmt.Sprintf("%s/%s", group.Name, d.Name)
	err = s.s3Client.Copy(s.c.Bucket, sourceKey, destinationKey)
	if err != nil {
		return err
	}

	d.Url = fmt.Sprintf("s3://%s/%s", s.c.Bucket, destinationKey)

	return s.repository.Create(d)
}

func (s service) FindById(id uint) (*model.Database, error) {
	d, err := s.repository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			idStr := strconv.FormatUint(uint64(id), 10)
			err = apperror.NewNotFound("database not found", idStr)
		}
	}
	return d, err
}

func (s service) Lock(id uint, instanceId uint, userId uint) (*model.Lock, error) {
	lock, err := s.repository.Lock(id, instanceId, userId)
	if err != nil {
		// TODO: Don't handle errors like this
		// Don't check the error by looking at the error message... Use the type
		// Don't allow gorm errors outside the repository
		if err.Error() == "record not found" {
			idStr := strconv.FormatUint(uint64(id), 10)
			err = apperror.NewNotFound("database not found", idStr)
		}

		// TODO: Don't handle errors like this
		// Don't check the error by looking at the error message... Use the type
		// Don't allow gorm errors outside the repository
		if strings.HasPrefix(err.Error(), "already locked by: ") {
			err = apperror.NewConflict(err.Error())
		}
	}
	return lock, err
}

func (s service) Unlock(id uint) error {
	err := s.repository.Unlock(id)
	if err != nil {
		if err.Error() == "record not found" {
			idStr := strconv.FormatUint(uint64(id), 10)
			err = apperror.NewNotFound("database not found", idStr)
		}
	}
	return err
}

func (s service) Upload(d *model.Database, group *models.Group, file io.Reader) (*model.Database, error) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(file)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s/%s", group.Name, d.Name)

	err = s.s3Client.Upload(s.c.Bucket, key, buffer)
	if err != nil {
		return nil, err
	}

	d.Url = fmt.Sprintf("s3://%s/%s", s.c.Bucket, key)

	err = s.repository.Save(d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (s service) Download(id uint, dst io.Writer, cb func(contentLength int64)) error {
	d, err := s.repository.FindById(id)
	if err != nil {
		return err
	}

	if d.Url == "" {
		return apperror.NewBadRequest(fmt.Sprintf("database with %d doesn't reference any url", id))
	}

	u, err := url.Parse(d.Url)
	if err != nil {
		return err
	}

	key := strings.TrimPrefix(u.Path, "/")
	return s.s3Client.Download(s.c.Bucket, key, dst, cb)
}

func (s service) Delete(id uint) error {
	d, err := s.repository.FindById(id)
	if err != nil {
		return err
	}

	u, err := url.Parse(d.Url)
	if err != nil {
		return err
	}

	key := strings.TrimPrefix(u.Path, "/")
	err = s.s3Client.Delete(s.c.Bucket, key)
	if err != nil {
		return err
	}

	return s.repository.Delete(id)
}

func (s service) List(groups []*models.Group) ([]*model.Database, error) {
	groupNames := make([]string, len(groups))
	for i, group := range groups {
		groupNames[i] = group.Name
	}

	instances, err := s.repository.FindByGroupNames(groupNames)
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (s service) Update(d *model.Database) error {
	return s.repository.Update(d)
}

func (s service) CreateExternalDownload(databaseID uint, expiration time.Time) (model.ExternalDownload, error) {
	err := s.repository.PurgeExternalDownload()
	if err != nil {
		return model.ExternalDownload{}, err
	}

	now := time.Now()
	if expiration.After(now) {
		return model.ExternalDownload{}, fmt.Errorf("expiration %s needs to be in the future (current %s)", expiration, now)
	}

	return s.repository.CreateExternalDownload(databaseID, expiration)
}

func (s service) FindExternalDownload(uuid uuid.UUID) (model.ExternalDownload, error) {
	err := s.repository.PurgeExternalDownload()
	if err != nil {
		return model.ExternalDownload{}, err
	}
	return s.repository.FindExternalDownload(uuid)
}
