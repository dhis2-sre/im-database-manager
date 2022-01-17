package database

import (
	"bytes"
	"fmt"
	"github.com/dhis2-sre/im-database-manager/internal/apperror"
	"github.com/dhis2-sre/im-database-manager/pkg/config"
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"github.com/dhis2-sre/im-database-manager/pkg/storage"
	jobClient "github.com/dhis2-sre/im-job/pkg/client"
	jobModels "github.com/dhis2-sre/im-job/swagger/sdk/models"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"io"
	"strconv"
	"strings"
)

type Service interface {
	Create(d *model.Database) error
	FindById(id uint) (*model.Database, error)
	Lock(id uint, instanceId uint) (*model.Database, error)
	Unlock(id uint) error
	Upload(d *model.Database, file io.Reader) (*model.Database, error)
	Delete(id uint) error
	List(groups []*models.Group) ([]*model.Database, error)
	Update(d *model.Database) error
	Save(token string, id uint) (string, error)
}

func ProvideService(c config.Config, s3Client storage.S3Client, jobClient jobClient.Client, repository Repository) Service {
	return &service{c, s3Client, jobClient, repository}
}

type service struct {
	c          config.Config
	s3Client   storage.S3Client
	jobClient  jobClient.Client
	repository Repository
}

func (s service) Create(d *model.Database) error {
	return s.repository.Create(d)
}

func (s service) FindById(id uint) (*model.Database, error) {
	d, err := s.repository.FindById(id)
	if err != nil {
		if err.Error() == "record not found" {
			idStr := strconv.FormatUint(uint64(id), 10)
			err = apperror.NewNotFound("database not found", idStr)
		}
	}
	return d, err
}

func (s service) Lock(id uint, instanceId uint) (*model.Database, error) {
	d, err := s.repository.Lock(id, instanceId)
	if err != nil {
		if err.Error() == "record not found" {
			idStr := strconv.FormatUint(uint64(id), 10)
			err = apperror.NewNotFound("database not found", idStr)
		}

		if strings.HasPrefix(err.Error(), "already locked by: ") {
			err = apperror.NewConflict(err.Error())
		}
	}
	return d, err
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

func (s service) Upload(d *model.Database, file io.Reader) (*model.Database, error) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(file)
	if err != nil {
		return nil, err
	}

	// TODO: Look up group name
	// TODO: Add file extension
	key := fmt.Sprintf("%d/%s", d.GroupID, d.Name)
	err = s.s3Client.Upload(s.c.Bucket, key, buffer)
	if err != nil {
		return nil, err
	}

	d.Url = fmt.Sprintf("s3://%s/%s", s.c.Bucket, key)
	return d, nil
}

func (s service) Delete(id uint) error {
	d, err := s.repository.FindById(id)
	if err != nil {
		return err
	}

	// TODO: Look up group name
	// TODO: Add file extension
	key := fmt.Sprintf("%d/%s", d.GroupID, d.Name)
	err = s.s3Client.Delete(s.c.Bucket, key)
	if err != nil {
		return err
	}

	return s.repository.Delete(id)
}

func (s service) List(groups []*models.Group) ([]*model.Database, error) {
	groupIds := make([]uint, len(groups))
	for i, group := range groups {
		groupIds[i] = uint(group.ID)
	}

	instances, err := s.repository.FindByGroupIds(groupIds)
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (s service) Update(d *model.Database) error {
	return s.repository.Update(d)
}

func (s service) Save(token string, id uint) (string, error) {
	d, err := s.FindById(id)
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("%d/%s", d.GroupID, d.Name)
	payload := map[string]string{
		"S3_BUCKET": s.c.Bucket,
		"S3_KEY":    key,
	}
	body := &jobModels.RunJobRequest{
		GroupID:  uint64(d.GroupID),
		Payload:  payload,
		TargetID: uint64(d.InstanceID),
	}

	runId, err := s.jobClient.Run(token, uint(3), body)
	if err != nil {
		return "", err
	}

	return runId, nil
}
