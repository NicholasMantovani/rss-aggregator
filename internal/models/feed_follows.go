package models

import (
	"time"

	"github.com/NicholasMantovani/rssaggregator/internal/database"
	"github.com/google/uuid"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
}

func DatabaseFeedFollowToFeedFollow(from database.FeedFollow) FeedFollow {
	return FeedFollow{ID: from.ID, CreatedAt: from.CreatedAt, UpdatedAt: from.UpdatedAt, UserId: from.UserID, FeedId: from.FeedID}
}

func DatabaseFeedFollowsToFeedFollows(from []database.FeedFollow) []FeedFollow {
	out := []FeedFollow{}
	for _, f := range from {
		out = append(out, DatabaseFeedFollowToFeedFollow(f))
	}
	return out
}
