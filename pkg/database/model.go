package database

import "gorm.io/gorm"

type Database struct {
	gorm.Model
	Name       string `gorm:"index:idx_name_and_group,unique"`
	GroupName  string `gorm:"index:idx_name_and_group,unique"`
	InstanceID uint   // instance which currently has the lock
	Url        string // s3... Path?
}
