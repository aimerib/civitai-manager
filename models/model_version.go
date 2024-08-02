package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
)

type ModelVersion struct {
	ID               int              `json:"-" db:"id"`
	CivitaiID        int              `json:"id" db:"civitai_id"`
	ModelID          int              `json:"-" db:"model_id"`
	Index            int              `json:"index" db:"index"`
	Name             string           `json:"name" db:"name"`
	BaseModel        string           `json:"baseModel" db:"base_model"`
	BaseModelType    string           `json:"baseModelType" db:"base_model_type"`
	PublishedAt      time.Time        `json:"publishedAt" db:"published_at"`
	Availability     string           `json:"availability" db:"availability"`
	NsfwLevel        int              `json:"nsfwLevel" db:"nsfw_level"`
	Description      *string          `json:"description" db:"description"`
	DownloadURL      *string          `json:"downloadUrl" db:"download_url"`
	TrainedWords     []string         `json:"trainedWords" db:"-"`
	TrainedWordsJSON string           `json:"-" db:"trained_words"`
	Stats            ModelVersionStat `json:"stats" has_one:"model_version_stat"`
	Files            Files            `json:"files" has_many:"files"`
	Images           Images           `json:"images" has_many:"images"`
	CreatedAt        time.Time        `json:"-" db:"created_at"`
	UpdatedAt        time.Time        `json:"-" db:"updated_at"`
}

type ModelVersions []ModelVersion

type TrainedWords []string

func (mv *ModelVersion) UnmarshalJSON(data []byte) error {
	type Alias ModelVersion
	aux := &struct {
		*Alias
		TrainedWords []string `json:"trainedWords"`
	}{
		Alias: (*Alias)(mv),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	mv.TrainedWords = aux.TrainedWords
	return nil
}

func (mv *ModelVersion) BeforeCreate(tx *pop.Connection) error {
	return mv.serializeTrainedWords()
}

func (mv *ModelVersion) BeforeUpdate(tx *pop.Connection) error {
	return mv.serializeTrainedWords()
}

func (mv *ModelVersion) AfterFind(tx *pop.Connection) error {
	return mv.deserializeTrainedWords()
}

func (mv *ModelVersion) serializeTrainedWords() error {
	if len(mv.TrainedWords) > 0 {
		jsonBytes, err := json.Marshal(mv.TrainedWords)
		if err != nil {
			return err
		}
		mv.TrainedWordsJSON = string(jsonBytes)
	}
	return nil
}

func (mv *ModelVersion) deserializeTrainedWords() error {
	if mv.TrainedWordsJSON != "" {
		return json.Unmarshal([]byte(mv.TrainedWordsJSON), &mv.TrainedWords)
	}
	return nil
}
