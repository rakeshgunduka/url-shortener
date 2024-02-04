package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name         string `gorm:"column:name"`
	UserAgent    string `gorm:"column:user_agent"`
	OSFamily     string `gorm:"column:os_family"`
	DeviceBrand  string `gorm:"column:device_brand"`
	DeviceFamily string `gorm:"column:device_family"`
	DeviceModel  string `gorm:"column:device_model"`
}
