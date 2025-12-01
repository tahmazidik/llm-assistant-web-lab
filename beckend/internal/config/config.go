package config

import (
	"os"
)

// HTTPConfig хранит настройки HTTP сервера
type HTTPConfig struct {
	Addr string
}

// Config - общая конфигурация приложения
// TODO: добавить другие настройки (БД, кэш, OpenAI и т.д.)
type Config struct {
	HTTP *HTTPConfig
}

// Load загружает конфигурацию из переменных окружения
func Load() *Config {
	addr := os.Getenv("APP_HTTP_ADDR")
	if addr == "" {
		addr = ":8080" // значение по умолчанию
	}

	return &Config{
		HTTP: &HTTPConfig{
			Addr: addr,
		},
	}
}
