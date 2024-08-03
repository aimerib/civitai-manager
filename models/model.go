package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

type Model struct {
	ID                    int             `json:"-" db:"id"`
	CivitaiID             int             `json:"id" db:"civitai_id"`
	Name                  *string         `json:"name" db:"name"`
	Description           *string         `json:"description" db:"description"`
	AllowNoCredit         bool            `json:"allowNoCredit" db:"allow_no_credit"`
	AllowDerivatives      bool            `json:"allowDerivatives" db:"allow_derivatives"`
	AllowDifferentLicense bool            `json:"allowDifferentLicense" db:"allow_different_license"`
	AllowCommericalUse    json.RawMessage `json:"allowCommercialUse" db:"allow_commercial_use"`
	Type                  string          `json:"type" db:"type"`
	Minor                 bool            `json:"minor" db:"minor"`
	Poi                   bool            `json:"poi" db:"poi"`
	Nsfw                  bool            `json:"nsfw" db:"nsfw"`
	NsfwLevel             int             `json:"nsfwLevel" db:"nsfw_level"`
	Cosmetic              nulls.String    `json:"cosmetic" db:"cosmetic"`
	CreatedAt             time.Time       `json:"-" db:"created_at"`
	UpdatedAt             time.Time       `json:"-" db:"updated_at"`
	ModelVersions         []ModelVersions `json:"modelVersions" has_many:"model_versions"`
	Stats                 Stat            `json:"stats" has_one:"stat"`
	Creator               Creator         `json:"creator" has_one:"creator"`
	Tags                  Tags            `json:"-" many_to_many:"model_tags" db:"-"`
	Checked               bool            `json:"-" db:"checked"`
}

type Models []Model

func (m *Model) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (m *Model) FindByCivitaiID(tx *pop.Connection, id int) error {
	return tx.Where("civitai_id = ?", id).First(m)
}

func (m *Model) UnmarshalJSON(data []byte) error {
	type Alias Model // Create an alias to avoid recursive UnmarshalJSON calls
	aux := &struct {
		*Alias
		Tags []string `json:"tags"`
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	m.Tags = make(Tags, len(aux.Tags))
	for i, tagName := range aux.Tags {
		m.Tags[i] = Tag{Name: &tagName}
	}
	return nil
}

func (m *Model) SaveWithTags(tx *pop.Connection) error {
	return tx.Transaction(func(tx *pop.Connection) error {
		// Save the model first
		if err := tx.Create(m); err != nil {
			return fmt.Errorf("error saving model: %w", err)
		}

		// Now handle the tags
		for _, tag := range m.Tags {
			if err := tag.FindOrCreate(tx); err != nil {
				return fmt.Errorf("error finding or creating tag: %w", err)
			}

			modelTag := ModelTag{
				ModelID: m.ID,
				TagID:   tag.ID,
			}
			if err := tx.Create(&modelTag); err != nil {
				return fmt.Errorf("error creating model_tag association: %w", err)
			}
		}

		return nil
	})
}
