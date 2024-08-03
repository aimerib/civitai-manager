package models

import (
	"encoding/json"
	"time"
)

type ModelVersions struct {
	ID            int              `json:"-" db:"id"`
	CivitaiID     int              `json:"id" db:"civitai_id"`
	ModelID       int              `json:"-" db:"model_id"`
	Index         int              `json:"index" db:"model_index"`
	Name          *string          `json:"name" db:"name"`
	BaseModel     *string          `json:"baseModel" db:"base_model"`
	BaseModelType *string          `json:"baseModelType" db:"base_model_type"`
	PublishedAt   time.Time        `json:"publishedAt" db:"published_at"`
	Availability  *string          `json:"availability" db:"availability"`
	NsfwLevel     int              `json:"nsfwLevel" db:"nsfw_level"`
	Description   *string          `json:"description" db:"description"`
	DownloadURL   string           `json:"downloadUrl" db:"download_url"`
	TrainedWords  json.RawMessage  `json:"trainedWords" db:"trained_words"`
	Stats         ModelVersionStat `json:"stats" has_one:"model_version_stat"`
	Files         Files            `json:"files" has_many:"files"`
	Images        Images           `json:"images" has_many:"images"`
	CreatedAt     time.Time        `json:"-" db:"created_at"`
	UpdatedAt     time.Time        `json:"-" db:"updated_at"`
	Model         Model            `json:"-" belongs_to:"model"`
}

// type ModelVersions []ModelVersions
