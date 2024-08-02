package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

type Stat struct {
	ID                int       `json:"-" db:"id"`
	ModelID           int       `json:"-" db:"model_id"`
	DownloadCount     int       `json:"downloadCount" db:"download_count"`
	FavoriteCount     int       `json:"favoriteCount" db:"favorite_count"`
	ThumbsUpCount     int       `json:"thumbsUpCount" db:"thumbs_up_count"`
	ThumbsDownCount   int       `json:"thumbsDownCount" db:"thumbs_down_count"`
	CommentCount      int       `json:"commentCount" db:"comment_count"`
	RatingCount       int       `json:"ratingCount" db:"rating_count"`
	Rating            float64   `json:"rating" db:"rating"`
	TippedAmountCount int       `json:"tippedAmountCount" db:"tipped_amount_count"`
	CreatedAt         time.Time `json:"-" db:"created_at"`
	UpdatedAt         time.Time `json:"-" db:"updated_at"`
}

type Stats []Stat

func (s *Stat) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
