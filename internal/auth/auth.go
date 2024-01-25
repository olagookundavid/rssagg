package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts an API key from the Authorization header of an HTTP request
// "Authorization: ApiKey <API key>"
func GetApiKey(headers http.Header) (string, error) {
	authHeaderValue := headers.Get("Authorization")
	if authHeaderValue == "" {
		return "", errors.New("no authentication info found")
	}

	authHeaderValueSplit := strings.Split(authHeaderValue, " ")
	if len(authHeaderValueSplit) != 2 {
		return "", errors.New("malformed auth header")
	}

	if authHeaderValueSplit[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return authHeaderValueSplit[1], nil
}
