package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

type Creator struct {
	ID        int       `db:"id"`
	Username  string    `json:"username" db:"username"`
	Image     *string   `json:"image" db:"image"`
	ModelID   int       `db:"model_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Creators []Creator

func (c *Creator) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
