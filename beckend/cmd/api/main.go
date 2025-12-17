package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/app"
	"github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/config"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	application := app.New(cfg)

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
