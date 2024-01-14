package models

import (
	"time"

	"github.com/NicholasMantovani/rssaggregator/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserId    uuid.UUID `json:"user_id"`
}

func DatabaseFeedToFeed(from database.Feed) Feed {
	return Feed{
		ID:        from.ID,
		CreatedAt: from.CreatedAt,
		UpdatedAt: from.UpdatedAt,
		Name:      from.Name,
		URL:       from.Url,
		UserId:    from.UserID,
	}

}

func DatabaseFeedsToFeeds(from []database.Feed) []Feed {
	feeds := []Feed{}

	for _, feed := range from {
		feeds = append(feeds, DatabaseFeedToFeed(feed))
	}

	return feeds
}
