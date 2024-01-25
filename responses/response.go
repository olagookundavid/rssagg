package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}

	// `json:"error"` is an struct tag
	// it will indicate to the json encoding library, how should we name this property
	// into json: {"error": "some error"}
	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errResponse{Error: msg})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to marshal JSON response")
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
