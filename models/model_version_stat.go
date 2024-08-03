package models

import "time"

type ModelVersionStat struct {
	ID              int       `json:"-" db:"id"`
	DownloadCount   int       `json:"downloadCount" db:"download_count"`
	RatingCount     int       `json:"ratingCount" db:"rating_count"`
	Rating          float64   `json:"rating" db:"rating"`
	ThumbsUpCount   int       `json:"thumbsUpCount" db:"thumbs_up_count"`
	ThumbsDownCount int       `json:"thumbsDownCount" db:"thumbs_down_count"`
	ModelVersionsID int       `json:"-" db:"model_versions_id"`
	CreatedAt       time.Time `jjson:"-" db:"created_at"`
	UpdatedAt       time.Time `json:"-" db:"updated_at"`
}

type ModelVersionStats []ModelVersionStat
