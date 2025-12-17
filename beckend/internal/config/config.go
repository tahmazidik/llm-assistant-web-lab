package config

import (
	"os"
)

// HTTPConfig хранит настройки HTTP сервера
type HTTPConfig struct {
	Addr string
}

// Config - общая конфигурация приложения
type Config struct {
	HTTP  *HTTPConfig
	DBDSN string

	OpenAIKey   string
	OpenAIModel string
}

// Load загружает конфигурацию из переменных окружения
func Load() *Config {
	addr := os.Getenv("APP_HTTP_ADDR")
	if addr == "" {
		addr = ":8080" // значение по умолчанию
	}

	dns := os.Getenv("DB_DSN")

	key := os.Getenv("OPENAI_API_KEY")
	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = "gpt-4o-mini"
	}

	return &Config{
		HTTP: &HTTPConfig{
			Addr: addr,
		},
		DBDSN:       dns,
		OpenAIKey:   key,
		OpenAIModel: model,
	}
}
