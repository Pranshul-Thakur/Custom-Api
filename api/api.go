package api

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type CoinbalanceParams struct { // Parameters that the API will take
	Username string `schema:"username"` // Made public and added schema tag
}

type CoinbalanceResponse struct { // Response with successful status code and account balance
	Code    int   `json:"code"`    // Response code - made public
	Balance int64 `json:"balance"` // Made public
}

type Error struct { // Response returned when error occurs
	Code    int    `json:"code"`    // Response code - made public
	Message string `json:"message"` // Made public
}

// Error handler functions
func RequestErrorHandler(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	response := Error{
		Code:    http.StatusBadRequest,
		Message: message,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
	}
}

func InternalErrorHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := Error{
		Code:    http.StatusInternalServerError,
		Message: "An unexpected error occurred",
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
	}
}

func UnauthorizedHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	response := Error{
		Code:    http.StatusUnauthorized,
		Message: "Not authorized",
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
	}
}
