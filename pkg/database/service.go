package database

import (
	"github.com/dhis2-sre/im-database-manager/pkg/model"
)

type Service interface {
	Create(d *model.Database) error
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
