package database

import (
	"github.com/dhis2-sre/im-database-manager/internal/apperror"
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"strconv"
	"strings"
)

type Service interface {
	Create(d *model.Database) error
	FindById(id uint) (*model.Database, error)
	Lock(id uint, instanceId uint) (*model.Database, error)
}

func ProvideService(repository Repository) Service {
	return &service{repository}
}

type service struct {
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
