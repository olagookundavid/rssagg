package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/olagookundavid/rssagg/internal/database"
	"github.com/olagookundavid/rssagg/models"
	"github.com/olagookundavid/rssagg/responses"
)

func (apiCfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, fmt.Sprint("Error parsing json:", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, fmt.Sprint("Error creating the feed:", err))
		return
	}

	responses.RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedToFeed(feed))
}

func (apiCfg *ApiConfig) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not get feeds: %v", err))
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, models.DatabaseFeedToFeeds(feeds))
}
