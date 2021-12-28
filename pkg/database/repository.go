package database

import (
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(d *model.Database) error
}

func ProvideRepository(DB *gorm.DB) Repository {
	return &repository{db: DB}
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(d *model.Database) error {
	return r.db.Create(&d).Error
}
