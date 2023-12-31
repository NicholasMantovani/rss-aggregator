package handlers

import (
	"net/http"

	"github.com/NicholasMantovani/rssaggregator/internal/utils"
)

func (a *ApiConfig) Readiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, 200, struct{}{})
}
