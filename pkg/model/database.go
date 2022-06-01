package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Database struct {
	gorm.Model
	Name              string `gorm:"index:idx_name_and_group,unique"`
	GroupName         string `gorm:"index:idx_name_and_group,unique"`
	InstanceID        uint   // instance which currently has the lock
	Url               string // s3... Path?
	ExternalDownloads []ExternalDownload
}

type ExternalDownload struct {
	UUID       uuid.UUID `gorm:"primaryKey;type:uuid"`
	Expiration time.Time
	DatabaseID uint
}
