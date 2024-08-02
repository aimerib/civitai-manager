package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

type ModelTag struct {
	ID        int       `db:"id"`
	ModelID   int       `db:"model_id"`
	TagID     int       `db:"tag_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ModelTags []ModelTag

func (mt *ModelTag) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
