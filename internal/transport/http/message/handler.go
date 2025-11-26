package message

import (
	"encoding/json"
	"log"
	"net/http"

	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/dialogs"
)

type Handler struct {
	DialogService dialogssvc.Service
}

type createMessageRequest struct {
	DialogID string `json:"dialog_id"`
	Sender   string `json:"sender"`
	Content  string `json:"content"`
}

func NewHandler(dialogService dialogssvc.Service) *Handler {
	return &Handler{
		DialogService: dialogService,
	}
}

// Create обрабатывает POST/messages запрос для создания нового сообщения
func (handler *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body createMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if body.DialogID == "" {
		http.Error(w, "dialog_id is required", http.StatusBadRequest)
		return
	}

	if body.Sender == "" {
		http.Error(w, "sender is required", http.StatusBadRequest)
		return
	}

	if body.Content == "" {
		http.Error(w, "content is required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Println("write error:", err)
	}
}

// List обрабатывает GET/messages запрос для получения списка сообщений
func (handler *Handler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
	if _, err := w.Write([]byte("message list not implemented")); err != nil {
		log.Println("write error:", err)
	}
}
