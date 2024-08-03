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
	Model     Model     `belongs_to:"model" db:"-"`
	Tag       Tag       `belongs_to:"tag" db:"-"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

type ModelTags []ModelTag

func (mt *ModelTag) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
