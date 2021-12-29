package database

import (
	"bytes"
	"fmt"
	"github.com/dhis2-sre/im-database-manager/internal/apperror"
	"github.com/dhis2-sre/im-database-manager/pkg/config"
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"github.com/dhis2-sre/im-database-manager/pkg/storage"
	"mime/multipart"
	"strconv"
	"strings"
)

type Service interface {
	Create(d *model.Database) error
	FindById(id uint) (*model.Database, error)
	Lock(id uint, instanceId uint) (*model.Database, error)
	Unlock(id uint) error
	Upload(d *model.Database, file *multipart.FileHeader) (*model.Database, error)
	Delete(id uint) error
}

func ProvideService(c config.Config, s3Client storage.S3Client, repository Repository) Service {
	return &service{c, s3Client, repository}
}

type service struct {
	c          config.Config
	s3Client   storage.S3Client
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

func (s service) Upload(d *model.Database, file *multipart.FileHeader) (*model.Database, error) {
	openFile, err := file.Open()
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(openFile)
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
