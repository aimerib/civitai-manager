package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

type Image struct {
	ID              int           `json:"-" db:"id"`
	ModelVersionsID int           `json:"-" db:"model_versions_id"`
	URL             string        `json:"url" db:"url"`
	NSFWLevel       int           `json:"nsfwLevel" db:"nsfw_level"`
	Width           int           `json:"width" db:"width"`
	Height          int           `json:"height" db:"height"`
	Hash            *string       `json:"hash" db:"hash"`
	Type            *string       `json:"type" db:"type"`
	HasMeta         bool          `json:"hasMeta" db:"has_meta"`
	OnSite          bool          `json:"onSite" db:"on_site"`
	CreatedAt       time.Time     `json:"-" db:"created_at"`
	UpdatedAt       time.Time     `json:"-" db:"updated_at"`
	ModelVersions   ModelVersions `json:"-" belongs_to:"model_versions"`
}

type Images []Image

func (i *Image) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
