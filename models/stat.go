package models

import (
	"time"

	"gorm.io/gorm"
)

type Stat struct {
	ID                uint      `gorm:"primaryKey" json:"-"`
	ModelID           uint      `gorm:"uniqueIndex" json:"-"`
	DownloadCount     int       `json:"downloadCount"`
	FavoriteCount     int       `json:"favoriteCount"`
	ThumbsUpCount     int       `json:"thumbsUpCount"`
	ThumbsDownCount   int       `json:"thumbsDownCount"`
	CommentCount      int       `json:"commentCount"`
	RatingCount       int       `json:"ratingCount"`
	Rating            float64   `json:"rating"`
	TippedAmountCount int       `json:"tippedAmountCount"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
}

func (s *Stat) BeforeCreate(tx *gorm.DB) (err error) {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return
}

func (s *Stat) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}
