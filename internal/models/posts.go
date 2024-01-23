package models

import (
	"time"

	"github.com/NicholasMantovani/rssaggregator/internal/database"
	"github.com/google/uuid"
)

type Post struct {
	ID               uuid.UUID `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Title            string    `json:"title"`
	Description      *string   `json:"description"`
	DescriptionEmpty string    `json:"description_empty,omitempty"` //If the string has its null value ("") this will not be present in the json
	PublishedAt      time.Time `json:"published_at"`
	Url              string    `json:"url"`
	FeedID           uuid.UUID `json:"feed_id"`
}

func DatabasePostToPost(from database.Post) Post {
	var description *string

	if from.Description.Valid {
		description = &from.Description.String
	}

	return Post{
		ID:               from.ID,
		CreatedAt:        from.CreatedAt,
		UpdatedAt:        from.UpdatedAt,
		Title:            from.Title,
		Description:      description,
		DescriptionEmpty: from.Description.String,
		Url:              from.Url,
		PublishedAt:      from.PublishedAt,
		FeedID:           from.FeedID,
	}
}

func DatabasePostsToPosts(from []database.Post) []Post {
	out := []Post{}
	for _, v := range from {
		out = append(out, DatabasePostToPost(v))
	}
	return out
}
