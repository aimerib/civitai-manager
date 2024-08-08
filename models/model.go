package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID                    uint           `gorm:"primaryKey" json:"-"`
	CivitaiID             int            `gorm:"uniqueIndex" json:"id"`
	Name                  string         `json:"name"`
	Description           string         `json:"description"`
	AllowNoCredit         bool           `json:"allowNoCredit"`
	AllowDerivatives      bool           `json:"allowDerivatives"`
	AllowDifferentLicense bool           `json:"allowDifferentLicense"`
	AllowCommericalUse    CommercialUse  `gorm:"type:json" json:"allowCommercialUse"`
	Type                  string         `json:"type"`
	Minor                 bool           `json:"minor"`
	Poi                   bool           `json:"poi"`
	Nsfw                  bool           `json:"nsfw"`
	NsfwLevel             int            `json:"nsfwLevel"`
	Cosmetic              string         `json:"cosmetic"`
	CreatedAt             time.Time      `json:"-"`
	UpdatedAt             time.Time      `json:"-"`
	ModelVersions         []ModelVersion `gorm:"foreignKey:ModelID" json:"modelVersions"`
	Stats                 Stat           `gorm:"foreignKey:ModelID" json:"stats"`
	Creator               Creator        `gorm:"foreignKey:ModelID" json:"creator"`
	Tags                  []Tag          `gorm:"many2many:model_tags;" json:"tags"`
	Checked               bool           `json:"-"`
}

type CommercialUse []string

func (cu *CommercialUse) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &cu)
}

func (cu CommercialUse) Value() (driver.Value, error) {
	return json.Marshal(cu)
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Model) FindByCivitaiID(db *gorm.DB, id int) error {
	return db.Where("civitai_id = ?", id).First(m).Error
}

func (m *Model) SaveWithTags(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(m).Error; err != nil {
			return err
		}

		for _, tag := range m.Tags {
			if err := tx.FirstOrCreate(&tag, Tag{Name: tag.Name}).Error; err != nil {
				return err
			}

			if err := tx.Model(m).Association("Tags").Append(&tag); err != nil {
				return err
			}
		}

		return nil
	})
}
