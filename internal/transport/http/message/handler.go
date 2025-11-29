package message

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/models"
	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/dialogs"
	httpresp "github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http/response"
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
		httpresp.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var body createMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpresp.Error(w, http.StatusBadRequest, "invalid json")
		return
	}

	if body.DialogID == "" {
		httpresp.Error(w, http.StatusBadRequest, "dialog_id is required")
		return
	}

	if body.Sender == "" {
		httpresp.Error(w, http.StatusBadRequest, "sender is required")
		return
	}

	sender := models.SenderType(body.Sender)
	if sender != models.SenderUser && sender != models.SenderAssistant {
		httpresp.Error(w, http.StatusBadRequest, "invalid sender type")
		return
	}

	ctx := r.Context()
	msg, err := handler.DialogService.AddMessage(
		ctx,
		models.DialogID(body.DialogID),
		sender,
		body.Content,
	)
	if err != nil {
		switch err {
		case dialogssvc.ErrEmptyMessageContent:
			httpresp.Error(w, http.StatusBadRequest, "message content is required")
		case dialogssvc.ErrDialogNotFound:
			httpresp.Error(w, http.StatusNotFound, "dialog not found")
		default:
			log.Println("add message error: ", err)
			httpresp.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	httpresp.JSON(w, http.StatusCreated, msg)
}

// List обрабатывает GET/messages запрос для получения списка сообщений
func (handler *Handler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpresp.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	dialogID := r.URL.Query().Get("dialog_id")
	if dialogID == "" {
		httpresp.Error(w, http.StatusBadRequest, "dialog_id is requierd")
		return
	}

	ctx := r.Context()
	msg, err := handler.DialogService.ListMessages(ctx, models.DialogID(dialogID))
	if err != nil {
		log.Println("list messages error: ", err)
		httpresp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	httpresp.JSON(w, http.StatusOK, msg)
}
