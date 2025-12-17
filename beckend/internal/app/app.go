package app

import (
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
	// Создаем репозиторий пользователей(in-memory)
	userRepo := memory2.NewUserRepository()
	// создаем сервис пользователей
	userService := userssvc.NewService(userRepo)

	var (
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

		_, err = db.Exec(`
			  INSERT INTO users (id, email, name, password_hash)
			  VALUES ($1, $2, $3, $4)
			  ON CONFLICT (email) DO NOTHING
			`,
			"00000000-0000-0000-0000-000000000001",
			"demo@local",
			"Demo User",
			"demo",
		)
		if err != nil {
			log.Fatal("seed demo user error:", err)
		}

		dialogRepo = postgres2.NewDialogRepository(db)
		messageRepo = postgres2.NewMessageRepository(db)
		log.Println("storage: postgres")
	} else {
		dialogRepo = memory2.NewDialogRepository()
		messageRepo = memory2.NewMessageRepository()
		log.Println("storage: in-memory")
	}

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
