package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/avukadin/goapi/api"
	"github.com/avukadin/goapi/internal/middleware/authorization"
	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

var decoder = schema.NewDecoder()

func Handle(r *chi.Mux) {
	// Apply authorization middleware
	r.Use(authorization.Middleware)

	// Define routes
	r.Route("/account", func(router chi.Router) {
		router.Get("/coins", GetCoinBalance)
	})
}

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinbalanceParams{}
	var decoder = schema.NewDecoder()

	// Parse query parameters
	err := decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Validate username parameter
	if params.Username == "" {
		log.Error("Username parameter is required")
		api.RequestErrorHandler(w, "Username parameter is required")
		return
	}

	// Get balance for user (mock implementation)
	balance := getCoinBalanceForUser(params.Username)

	// Create response
	response := api.CoinbalanceResponse{
		Code:    http.StatusOK,
		Balance: balance,
	}

	// Set content type and write response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

// Mock function to simulate getting coin balance for a user
func getCoinBalanceForUser(username string) int64 {
	// In a real implementation, this would query a database
	// For now, return a mock balance based on username
	switch username {
	case "admin":
		return 10000
	case "user1":
		return 5000
	case "user2":
		return 2500
	default:
		return 1000
	}
}
