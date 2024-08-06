package models

import (
	"time"

	"gorm.io/gorm"
)

type ModelVersionStat struct {
	ID              uint      `gorm:"primaryKey" json:"-"`
	DownloadCount   int       `json:"downloadCount"`
	RatingCount     int       `json:"ratingCount"`
	Rating          float64   `json:"rating"`
	ThumbsUpCount   int       `json:"thumbsUpCount"`
	ThumbsDownCount int       `json:"thumbsDownCount"`
	ModelVersionID  uint      `gorm:"uniqueIndex" json:"-"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`

	// ModelVersion ModelVersion `gorm:"foreignKey:ModelVersionID" json:"-"`
}

func (mvs *ModelVersionStat) BeforeCreate(tx *gorm.DB) (err error) {
	mvs.CreatedAt = time.Time{}
	mvs.UpdatedAt = time.Now()
	return
}

func (mvs *ModelVersionStat) BeforeUpdate(tx *gorm.DB) (err error) {
	mvs.UpdatedAt = time.Now()
	return
}
