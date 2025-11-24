package dialogs

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/models"
	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/dialogs"
)

// Handler отвечает за HTTP-операции, связанные с диалогами
type Handler struct {
	DialogService dialogssvc.Service
}

// NewHandler создает новый Handler для диалогов
func NewHandler(dialogService dialogssvc.Service) *Handler {
	return &Handler{
		DialogService: dialogService,
	}
}

type createdDialogResponse struct {
	UserID string `json:"user_id"`
	Title  string `json:"title"`
}

// Create обрабатывает POST/dialogs запрос для создания нового диалога
func (handler *Handler) Create(w http.ResponseWriter, r *http.Request) {
	// Разрешаем только POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Парсим JSON из тела запроса
	var body createdDialogResponse
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// Проверка полей, что они не пустые
	if body.UserID == "" {
		http.Error(w, "user_id required", http.StatusBadRequest)
		return
	}

	if body.Title == "" {
		http.Error(w, "title is requierd", http.StatusBadRequest)
		return
	}

	// Вызываем бизнес-логику для создания диалога
	ctx := r.Context()
	dialog, err := handler.DialogService.CreateDialog(
		ctx,
		models.UserID(body.UserID),
		body.Title,
	)

	if err != nil {
		switch err {
		case dialogssvc.ErrEmptyDialogTitle:
			http.Error(w, "title cannot be empty", http.StatusBadRequest)
		default:
			log.Println("create dialog error: ", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(dialog); err != nil {
		log.Println("encode response error: ", err)
	}
}
