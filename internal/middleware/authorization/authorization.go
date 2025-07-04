package authorization

import (
	"net/http"

	"github.com/avukadin/goapi/api"
	log "github.com/sirupsen/logrus"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get authorization header
		authHeader := r.Header.Get("Authorization")

		// Check if authorization header is present
		if authHeader == "" {
			log.Error("Authorization header is missing")
			api.RequestErrorHandler(w, "Authorization header is required")
			return
		}

		// Simple token validation (in production, use proper JWT validation)
		if !isValidToken(authHeader) {
			log.Error("Invalid authorization token")
			api.UnauthorizedHandler(w)
			return
		}

		// If authorized, continue to next handler
		next.ServeHTTP(w, r)
	})
}

// Simple token validation function
// In production, implement proper JWT token validation
func isValidToken(token string) bool {
	// Remove "Bearer " prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// Mock token validation - in production, validate against JWT or database
	validTokens := map[string]bool{
		"admin-token": true,
		"user1-token": true,
		"user2-token": true,
		"valid-token": true,
	}

	return validTokens[token]
}
