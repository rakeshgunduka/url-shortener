package models

import "gorm.io/gorm"

type TokenRange struct {
	gorm.Model
	Start   int64  `gorm:"column:start"`
	End     int64  `gorm:"column:end"`
	Service string `gorm:"column:service"`
}
