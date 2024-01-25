package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/olagookundavid/rssagg/internal/database"
	"github.com/olagookundavid/rssagg/models"
	"github.com/olagookundavid/rssagg/responses"
)

func (apiCfg *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, fmt.Sprint("Error parsing json:", err))
		return
	}

	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedId,
		UserID:    user.ID,
	})

	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error following the feed: %v", err))
		return
	}

	responses.RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedFollowToFeedFollow(feed_follow))
}

func (apiCfg *ApiConfig) GetUserFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollowsByUserId(r.Context(), user.ID)

	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not get feed follows: %v", err))
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, models.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *ApiConfig) DeleteUserFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowId, err := uuid.Parse(chi.URLParam(r, "feedFollowId"))
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not parse feedFollowId into a valid uuid: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not delete feed follows: %v", err))
		return
	}

	responses.RespondWithJSON(w, http.StatusNoContent, struct{}{})
}
