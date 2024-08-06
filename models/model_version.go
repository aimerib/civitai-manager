package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type ModelVersion struct {
	ID            uint      `gorm:"primaryKey" json:"-"`
	CivitaiID     int       `gorm:"uniqueIndex" json:"id"`
	ModelID       uint      `json:"-"`
	Index         int       `json:"index"`
	Name          *string   `json:"name"`
	BaseModel     *string   `json:"baseModel"`
	BaseModelType *string   `json:"baseModelType"`
	PublishedAt   time.Time `json:"publishedAt"`
	Availability  *string   `json:"availability"`
	NsfwLevel     int       `json:"nsfwLevel"`
	Description   *string   `json:"description"`
	DownloadURL   string    `json:"downloadUrl"`
	TrainedWords  Words     `gorm:"type:json" json:"trainedWords"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`

	Model            Model            `gorm:"foreignKey:ModelID" json:"-"`
	ModelVersionStat ModelVersionStat `gorm:"foreignKey:ModelVersionID" json:"stats"`
	Files            []File           `gorm:"foreignKey:ModelVersionID" json:"files"`
	Images           []Image          `gorm:"foreignKey:ModelVersionID" json:"images"`
}

type Words []string

func (w *Words) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &w)
}

func (w Words) Value() (driver.Value, error) {
	return json.Marshal(w)
}

func (mv *ModelVersion) BeforeCreate(tx *gorm.DB) (err error) {
	mv.CreatedAt = time.Now()
	mv.UpdatedAt = time.Now()
	return
}

func (mv *ModelVersion) BeforeUpdate(tx *gorm.DB) (err error) {
	mv.UpdatedAt = time.Now()
	return
}
