package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/NicholasMantovani/rssaggregator/internal/auth"
	"github.com/NicholasMantovani/rssaggregator/internal/database"
	"github.com/NicholasMantovani/rssaggregator/internal/models"
	"github.com/NicholasMantovani/rssaggregator/internal/utils"
	"github.com/google/uuid"
)

func (a *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	}

	utils.RespondWithJson(w, 201, models.DatabaseUserToUser(user))
}

func (a *ApiConfig) GetUser(w http.ResponseWriter, r *http.Request) {

	apiKey, err := auth.GetApiKey(w.Header())
	if err != nil {
		utils.RespondWithError(w, 401, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := a.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not get user: %v", err))
	}

	utils.RespondWithJson(w, 200, models.DatabaseUserToUser(user))
}
