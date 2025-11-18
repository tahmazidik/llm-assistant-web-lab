package app

import (
	"log"
	"net/http"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/config"
	apphttp "github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http"
)

// Application хранит все, что нужно для запуска сервера
type Application struct {
	cfg    *config.Config //на каком адресе запускать сервер
	router *http.ServeMux // HTTP роутер
}

// New создает новое приложение: настраивает роутер и порт
func New(cfg *config.Config) *Application {
	// создаем роутер
	router := apphttp.NewRouter()

	// кладем все в Application
	return &Application{
		cfg:    cfg,
		router: router,
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
