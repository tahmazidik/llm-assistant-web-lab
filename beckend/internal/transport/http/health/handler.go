package health

import (
	"log"
	"net/http"
)

// Handler — простой обработчик для /health.
func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("ok")); err != nil {
		log.Println("write response error:", err)
	}
}
