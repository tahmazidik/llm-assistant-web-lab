package app

import (
	"log"
	"net/http"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/config"
	apphttp "github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http"

	memory "github.com/tahmazidik/llm-assistant-web-lab/internal/database/memory"
	usersvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/users"
)

// Application хранит все, что нужно для запуска сервера
type Application struct {
	cfg         *config.Config //на каком адресе запускать сервер
	router      *http.ServeMux // HTTP роутер
	userService usersvc.Service
}

// New создает новое приложение: настраивает роутер и порт
func New(cfg *config.Config) *Application {
	// Создаем репозиторий пользователей(in-memory)
	userRepo := memory.NewUserRepository()
	// создаем сервис пользователей
	userService := usersvc.NewService(userRepo)
	// создаем роутер
	router := apphttp.NewRouter()

	// кладем все в Application
	return &Application{
		cfg:         cfg,
		router:      router,
		userService: userService,
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
