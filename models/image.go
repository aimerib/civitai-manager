package models

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID             uint      `gorm:"primaryKey" json:"-"`
	ModelVersionID uint      `gorm:"index" json:"-"`
	URL            string    `json:"url"`
	NSFWLevel      int       `json:"nsfwLevel"`
	Width          int       `json:"width"`
	Height         int       `json:"height"`
	Hash           string    `json:"hash"`
	Type           string    `json:"type"`
	HasMeta        bool      `json:"hasMeta"`
	OnSite         bool      `json:"onSite"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

func (i *Image) BeforeCreate(tx *gorm.DB) (err error) {
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()
	return
}

func (i *Image) BeforeUpdate(tx *gorm.DB) (err error) {
	i.UpdatedAt = time.Now()
	return
}
