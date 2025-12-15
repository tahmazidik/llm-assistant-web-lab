package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/models"
)

var (
	ErrNotAuthHeader = errors.New("authorization header is missing")
	ErrInvalidToken  = errors.New("invalid auth token")
)

const demoUserID = models.UserID("00000000-0000-0000-0000-000000000001")
const demoToken = "demo-token-123"

// UserIDFromRequest достает user_id из заголовка Authorization
func UserIDFromRequest(r *http.Request) (models.UserID, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNotAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", ErrInvalidToken
	}

	token := parts[1]
	if token != demoToken {
		return "", ErrInvalidToken
	}

	return demoUserID, nil
}
