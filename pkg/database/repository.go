package database

import (
	"fmt"

	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(d *model.Database) error
	Save(d *model.Database) error
	FindById(id uint) (*model.Database, error)
	Lock(id uint, instanceId uint) (*model.Database, error)
	Unlock(id uint) error
	Delete(id uint) error
	FindByGroupNames(names []string) ([]*model.Database, error)
	Update(d *model.Database) error
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

func (r repository) Save(d *model.Database) error {
	return r.db.Save(&d).Error
}

func (r repository) FindById(id uint) (*model.Database, error) {
	var d *model.Database
	err := r.db.First(&d, id).Error
	return d, err
}

func (r repository) Lock(id uint, instanceId uint) (*model.Database, error) {
	var d *model.Database

	errTx := r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&d, id).Error
		if err != nil {
			return err
		}

		if d.InstanceID != 0 {
			return fmt.Errorf("already locked by: %d", d.InstanceID)
		}

		err = tx.Model(&d).Update("instance_id", instanceId).Error
		if err != nil {
			return err
		}

		return nil
	})

	return d, errTx
}

func (r repository) Unlock(id uint) error {
	err := r.db.Model(&model.Database{
		Model: gorm.Model{
			ID: id,
		},
	}).Update("instance_id", 0).Error

	if err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(id uint) error {
	return r.db.Unscoped().Delete(&model.Database{}, id).Error
}

func (r repository) FindByGroupNames(names []string) ([]*model.Database, error) {
	var databases []*model.Database

	err := r.db.
		Where("group_name IN ?", names).
		Find(&databases).Error

	return databases, err
}

func (r repository) Update(d *model.Database) error {
	return r.db.Save(d).Error
}
