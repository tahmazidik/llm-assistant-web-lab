package http

import (
	"log"
	stdhttp "net/http"
)

// NewRouter создает и настраивает HTTP-роутер приложения
func NewRouter() *stdhttp.ServeMux {
	mux := stdhttp.NewServeMux()

	// эндпоинт для проверки живости сервера
	mux.HandleFunc("/health", HealthHandler)

	return mux
}

func HealthHandler(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	w.WriteHeader(stdhttp.StatusOK)
	if _, err := w.Write([]byte("ok")); err != nil {
		log.Println("write response error:", err)
	}
}
