package handlers

import (
	"net/http"

	"github.com/olagookundavid/rssagg/responses"
)

// function signature that you need to use of you want to define a HTTP handler
// from interface http.Handler
func HandlerErr(w http.ResponseWriter, r *http.Request) {
	responses.RespondWithError(w, 500, "Something went wrong")
}
