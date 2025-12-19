package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/config"
	memory2 "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/database/memory"
	postgres2 "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/database/postgres"
	assistantsvc "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/services/assistant"
	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/services/dialogs"
	userssvc "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/services/users"
	apphttp "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/transport/http"

	_ "github.com/lib/pq"
)

// Application хранит все, что нужно для запуска сервера
type Application struct {
	cfg              *config.Config // на каком адресе запускать сервер
	router           *http.ServeMux // HTTP роутер
	userService      userssvc.Service
	dialogService    dialogssvc.Service
	assistantService assistantsvc.Service
}

// New создает новое приложение: настраивает роутер и порт
func New(cfg *config.Config) *Application {
	var (
		userRepo    userssvc.Repository
		dialogRepo  dialogssvc.DialogRepository
		messageRepo dialogssvc.MessageRepository
	)

	if cfg.DBDSN != "" {
		db, err := sql.Open("postgres", cfg.DBDSN)
		if err != nil {
			log.Fatal("db open error:", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatal("db ping error:", err)
		}

		userRepo = postgres2.NewUserRepository(db)
		dialogRepo = postgres2.NewDialogRepository(db)
		messageRepo = postgres2.NewMessageRepository(db)
		log.Println("storage: postgres")
	} else {
		userRepo = memory2.NewUserRepository()
		dialogRepo = memory2.NewDialogRepository()
		messageRepo = memory2.NewMessageRepository()
		log.Println("storage: in-memory")
	}

	userService := userssvc.NewService(userRepo)
	dialogService := dialogssvc.NewService(dialogRepo, messageRepo)

	var assistantService assistantsvc.Service

	if cfg.OpenAIKey == "" {
		log.Println("openai: OPENAI_API_KEY is empty (assistant replace disabled)")
	} else {
		svc, err := assistantsvc.New(dialogService, assistantsvc.Config{
			APIKey:      cfg.OpenAIKey,
			MaxHistory:  20,
			SystemIntro: "You are a helpful assistant. Answer clearly and concisely.",
		})

		if err != nil {
			log.Println("openai: assistant service init error:", err)
		} else {
			assistantService = svc
		}
	}

	// Пытаемся создать демо-пользователя, если его нет
	if _, err := userService.Register(context.Background(), "demo@local", "demo", "Demo User"); err != nil && err != userssvc.ErrEmailAlreadyInUse {
		log.Println("seed demo user error:", err)
	}

	// создаем роутер
	router := apphttp.NewRouter(userService, dialogService, assistantService)

	// кладем все в Application
	return &Application{
		cfg:              cfg,
		router:           router,
		userService:      userService,
		dialogService:    dialogService,
		assistantService: assistantService,
	}
}

// Run запускает HTTP сервер
func (app *Application) Run() error {
	addr := app.cfg.HTTP.Addr
	// пишем что стартуем
	log.Println("starting http server on", addr)
	// запускаем
	return http.ListenAndServe(addr, apphttp.WithCORS(app.router))
}
