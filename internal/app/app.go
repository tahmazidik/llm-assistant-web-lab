package app

import (
	"log"
	"net/http"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/config"
	apphttp "github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http"

	memory "github.com/tahmazidik/llm-assistant-web-lab/internal/database/memory"
	userssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/users"

	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/dialogs"
)

// Application хранит все, что нужно для запуска сервера
type Application struct {
	cfg           *config.Config // на каком адресе запускать сервер
	router        *http.ServeMux // HTTP роутер
	userService   userssvc.Service
	dialogService dialogssvc.Service
}

// New создает новое приложение: настраивает роутер и порт
func New(cfg *config.Config) *Application {
	// Создаем репозиторий пользователей(in-memory)
	userRepo := memory.NewUserRepository()
	// создаем сервис пользователей
	userService := userssvc.NewService(userRepo)

	// Репозитории для диалогов и сообщений(in-memory)
	dialogRepo := memory.NewDialogRepository()
	messageRepo := memory.NewMessageRepository()
	dialogService := dialogssvc.NewService(dialogRepo, messageRepo)

	// создаем роутер
	router := apphttp.NewRouter(userService, dialogService)

	// кладем все в Application
	return &Application{
		cfg:           cfg,
		router:        router,
		userService:   userService,
		dialogService: dialogService,
	}
}

// Run запускает HTTP сервер
func (app *Application) Run() error {
	addr := app.cfg.HTTP.Addr
	// пишем что стартуем
	log.Println("starting http server on", addr)
	// запускаем
	return http.ListenAndServe(addr, app.router)
}
