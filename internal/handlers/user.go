package handlers

import (
	"fmt"
	"net/http"

	"github.com/NicholasMantovani/rssaggregator/internal/utils"
)

func (a *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, 200, struct{}{})
	fmt.Print(a)
}
