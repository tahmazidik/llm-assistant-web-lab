package main

import (
	"log"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/app"
	"github.com/tahmazidik/llm-assistant-web-lab/internal/config"
)

func main() {
	cfg := config.Load()

	application := app.New(cfg)

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
