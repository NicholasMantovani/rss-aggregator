package handlers

import (
	"net/http"

	"github.com/NicholasMantovani/rssaggregator/internal/utils"
)

func Error(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 400, "Something went wrong")
}
