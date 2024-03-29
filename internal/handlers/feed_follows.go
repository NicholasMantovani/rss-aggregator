package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/NicholasMantovani/rssaggregator/internal/database"
	"github.com/NicholasMantovani/rssaggregator/internal/models"
	"github.com/NicholasMantovani/rssaggregator/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (a *ApiConfig) HandleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := parameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := a.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not create feed follow: %v", err))
		return
	}

	utils.RespondWithJson(w, 201, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (a *ApiConfig) HandleGetFeedsFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := a.DB.GetFeedsFollow(r.Context(), user.ID)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not get feed follows: %v", err))
		return
	}

	utils.RespondWithJson(w, 201, models.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (a *ApiConfig) HandleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	id := chi.URLParam(r, "id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not parse feed follow id: %v", err))
		return
	}

	err = a.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{ID: uuid, UserID: user.ID})
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Could not delete feed follow: %v", err))
		return
	}

	utils.RespondWithJson(w, http.StatusOK, struct{}{})

}
