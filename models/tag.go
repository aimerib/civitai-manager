package models

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	Name      *string   `gorm:"uniqueIndex" json:"name"`
	Models    []Model   `gorm:"many2many:model_tags;" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Tags []Tag

func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return
}

func (t *Tag) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return
}

func (t *Tag) FindOrCreate(db *gorm.DB) error {
	result := db.Where("name = ?", t.Name).FirstOrCreate(t)
	return result.Error
}
