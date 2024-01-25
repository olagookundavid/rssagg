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

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, fmt.Sprint("Error parsing json:", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, fmt.Sprint("Error creating the user: ", err))
		return
	}
	//databaseUserToUser(user) so json returned can be formatted well as lower camel case
	responses.RespondWithJSON(w, http.StatusCreated, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	responses.RespondWithJSON(w, http.StatusOK, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandleGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not get posts: %v", err))
	}

	responses.RespondWithJSON(w, http.StatusOK, models.DatabasePostToPosts(posts))
}
