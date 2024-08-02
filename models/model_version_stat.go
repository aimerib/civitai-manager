package models

import "time"

type ModelVersionStat struct {
	ID              int       `json:"id" db:"id"`
	DownloadCount   int       `json:"civitai_id" db:"civitai_id"`
	RatingCount     int       `json:"model_id" db:"model_id"`
	Rating          int       `json:"index" db:"index"`
	ThumbsUpCount   string    `json:"name" db:"name"`
	ThumbsDownCount string    `json:"base_model" db:"base_model"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type ModelVersionStats []ModelVersionStat
