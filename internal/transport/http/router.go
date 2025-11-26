package http

import (
	stdhttp "net/http"

	dialogshttp "github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http/dialogs"
	"github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http/health"

	usershttp "github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http/users"

	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/dialogs"
	userssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/users"
)

// NewRouter создает и настраивает HTTP-роутер приложения
func NewRouter(userService userssvc.Service, dialogService dialogssvc.Service) *stdhttp.ServeMux {
	mux := stdhttp.NewServeMux()

	// эндпоинт для проверки живости сервера
	mux.HandleFunc("/health", health.Handler)

	// хендлер пользователей
	userHandler := usershttp.NewUserHandler(userService)
	// хендлер диалогов
	dialogHandler := dialogshttp.NewHandler(dialogService)

	mux.HandleFunc("/users/register", userHandler.Register)
	mux.HandleFunc("/users/login", userHandler.Login)
	mux.HandleFunc("/dialogs", dialogHandler.Create)
	mux.HandleFunc("/dialogs/list", dialogHandler.List)

	return mux
}
