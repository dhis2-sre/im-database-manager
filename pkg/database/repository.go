package database

import (
	"fmt"
	"gorm.io/gorm"
)

type Repository interface {
	Create(d *Database) error
	Save(d *Database) error
	FindById(id uint) (*Database, error)
	Lock(id uint, instanceId uint) (*Database, error)
	Unlock(id uint) error
	Delete(id uint) error
	FindByGroupNames(names []string) ([]*Database, error)
	Update(d *Database) error
}

func ProvideRepository(DB *gorm.DB) Repository {
	return &repository{db: DB}
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(d *Database) error {
	return r.db.Create(&d).Error
}

func (r repository) Save(d *Database) error {
	return r.db.Save(&d).Error
}

func (r repository) FindById(id uint) (*Database, error) {
	var d *Database
	err := r.db.First(&d, id).Error
	return d, err
}

func (r repository) Lock(id uint, instanceId uint) (*Database, error) {
	var d *Database

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
	err := r.db.Model(&Database{
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
	return r.db.Unscoped().Delete(&Database{}, id).Error
}

func (r repository) FindByGroupNames(names []string) ([]*Database, error) {
	var databases []*Database

	err := r.db.
		Where("group_name IN ?", names).
		Find(&databases).Error

	return databases, err
}

func (r repository) Update(d *Database) error {
	return r.db.Save(d).Error
}
