package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

type Tag struct {
	ID        int       `json:"-" db:"id"`
	Name      string    `json:"name" db:"name"`
	Models    Models    `many_to_many:"model_tags"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

type Tags []Tag

func (t *Tag) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (t *Tag) FindOrCreate(tx *pop.Connection) error {
	existing := &Tag{}
	err := tx.Where("name = ?", t.Name).First(existing)
	if err == nil {
		*t = *existing
		return nil
	}
	return tx.Create(t)
}
