package model

import "gorm.io/gorm"

type Database struct {
	gorm.Model
	Name       string `gorm:"index;unique"`
	GroupId    uint
	InstanceID uint   // instance which currently has the lock
	Format     string // sql or pgc... Probably not necessary
	Url        string // s3... Path?
}
