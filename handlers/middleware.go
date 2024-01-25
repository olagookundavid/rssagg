package handlers

import (
	"fmt"
	"net/http"

	"github.com/olagookundavid/rssagg/internal/auth"
	"github.com/olagookundavid/rssagg/internal/database"
	"github.com/olagookundavid/rssagg/responses"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			responses.RespondWithError(w, http.StatusForbidden, fmt.Sprint("Error parsing json:", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			responses.RespondWithError(w, http.StatusNotFound, fmt.Sprint("Could not get the user:", err))
			return
		}

		handler(w, r, user)
	}
}
