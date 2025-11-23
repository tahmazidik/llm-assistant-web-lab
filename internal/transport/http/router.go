package http

import (
	stdhttp "net/http"

	userssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/users"
	"github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http/health"
	usershttp "github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http/users"
)

// NewRouter создает и настраивает HTTP-роутер приложения
func NewRouter(userService userssvc.Service) *stdhttp.ServeMux {
	mux := stdhttp.NewServeMux()

	// эндпоинт для проверки живости сервера
	mux.HandleFunc("/health", health.Handler)

	// хендлер пользователей
	userHandler := usershttp.NewUserHandler(userService)
	mux.HandleFunc("/users/register", userHandler.Register)
	mux.HandleFunc("/users/login", userHandler.Login)

	return mux
}
