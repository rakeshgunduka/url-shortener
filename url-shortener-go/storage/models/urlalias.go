package models

import (
	"time"

	"gorm.io/gorm"
)

type UrlAliasMap struct {
	gorm.Model
	ID      int       `gorm:"primaryKey;autoIncrement"`
	Alias   string    `gorm:"column:alias"`
	URL     string    `gorm:"column:url"`
	Created time.Time `gorm:"column:created"`
	Expiry  time.Time `gorm:"column:expiry"`
}
