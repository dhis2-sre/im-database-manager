package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"

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
	CreateExternalDownload(databaseID uint, expiration time.Time) (model.ExternalDownload, error)
	FindExternalDownload(uuid uuid.UUID) (model.ExternalDownload, error)
	PurgeExternalDownload() error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(DB *gorm.DB) *repository {
	return &repository{db: DB}
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

func (r repository) CreateExternalDownload(databaseID uint, expiration time.Time) (model.ExternalDownload, error) {
	externalDownload := model.ExternalDownload{
		UUID:       uuid.New(),
		Expiration: expiration,
		DatabaseID: databaseID,
	}

	err := r.db.Save(externalDownload).Error

	return externalDownload, err
}

func (r repository) FindExternalDownload(uuid uuid.UUID) (model.ExternalDownload, error) {
	var d model.ExternalDownload
	err := r.db.
		Where("expiration > ?", time.Now().UTC()).
		First(&d, uuid).Error
	return d, err
}

func (r repository) PurgeExternalDownload() error {
	var d model.ExternalDownload
	err := r.db.
		Where("expiration < ?", time.Now().UTC()).
		Delete(&d).Error
	return err
}
