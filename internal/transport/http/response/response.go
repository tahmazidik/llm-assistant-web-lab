package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// errorResponse представляет структуру ошибки в ответе
type errorResponse struct {
	Error string `json:"error"`
}

// JSON отправляет произвольный объект как JSON с нужным статусом
func JSON(w http.ResponseWriter, status int, resp any) {
	w.Header().Set("Content-Type", "application/json; chraset=utf-8")
	w.WriteHeader(status)

	if resp == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("write json response error:", err)
	}
}

// Error отправляет ошибку JSON-ошибку в виде {"error": "message"}
func Error(w http.ResponseWriter, status int, message string) {
	resp := errorResponse{
		Error: message,
	}
	JSON(w, status, resp)
}
