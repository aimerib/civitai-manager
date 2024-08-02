package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

type ModelTag struct {
	ID        int       `json:"id" db:"id"`
	ModelID   int       `json:"model_id" db:"model_id"`
	TagID     int       `json:"tag_id" db:"tag_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ModelTags []ModelTag

func (mt *ModelTag) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
