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

func (a *ApiConfig) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := a.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Cloud not create user: %v", err))
		return
	}

	utils.RespondWithJson(w, 201, models.DatabaseUserToUser(user))
}

func (a *ApiConfig) HandleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJson(w, 200, models.DatabaseUserToUser(user))
}

func (a *ApiConfig) HandleGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := a.DB.GetPostsForUsers(r.Context(), database.GetPostsForUsersParams{UserID: user.ID, Limit: 10})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Cloud not get posts from user: %v", err))
		return
	}

	utils.RespondWithJson(w, http.StatusOK, models.DatabasePostsToPosts(posts))
}
