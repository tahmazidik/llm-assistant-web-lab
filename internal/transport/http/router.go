package http

import (
	stdhttp "net/http"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http/health"

	dialogshttp "github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http/dialogs"
	messagehttp "github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http/message"
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
	mux.HandleFunc("/users/register", userHandler.Register)
	mux.HandleFunc("/users/login", userHandler.Login)

	// хендлер диалогов
	dialogHandler := dialogshttp.NewHandler(dialogService)
	mux.HandleFunc("/dialogs", dialogHandler.Create)
	mux.HandleFunc("/dialogs/list", dialogHandler.List)

	// хендлер сообщений
	messageHandler := messagehttp.NewHandler(dialogService)
	mux.HandleFunc("/messages", messageHandler.Create)
	mux.HandleFunc("/messages/list", messageHandler.List)

	return mux
}
