package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/NicholasMantovani/rssaggregator/internal/database"
	"github.com/NicholasMantovani/rssaggregator/internal/models"
	"github.com/NicholasMantovani/rssaggregator/internal/utils"
	"github.com/google/uuid"
)

func (a *ApiConfig) HandleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params := parameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := a.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Cloud not create feed: %v", err))
		return
	}

	utils.RespondWithJson(w, 201, models.DatabaseFeedToFeed(feed))
}

func (a *ApiConfig) HandleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := a.DB.GetFeeds(r.Context())

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Cloud not create user: %v", err))
		return
	}

	utils.RespondWithJson(w, 201, models.DatabaseFeedsToFeeds(feeds))
}
