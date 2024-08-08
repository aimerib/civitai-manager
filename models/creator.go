package models

import (
	"time"

	"gorm.io/gorm"
)

type Creator struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	Username  string    `json:"username"`
	Image     string    `json:"image"`
	ModelID   uint      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Model Model `gorm:"foreignKey:ModelID" json:"-"`
}

func (c *Creator) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return
}

func (c *Creator) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}
