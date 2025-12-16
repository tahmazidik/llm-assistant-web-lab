package message

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	models2 "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/models"
	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/services/dialogs"
	authhttp "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/transport/http/auth"
	httpresp "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/transport/http/response"
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

	_, err := authhttp.UserIDFromRequest(r)
	if err != nil {
		if errors.Is(err, authhttp.ErrNotAuthHeader) || errors.Is(err, authhttp.ErrInvalidToken) {
			httpresp.Error(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		log.Println("auth error: ", err)
		httpresp.Error(w, http.StatusInternalServerError, "internal server error")
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

	sender := models2.SenderType(body.Sender)
	if sender != models2.SenderUser && sender != models2.SenderAssistant {
		httpresp.Error(w, http.StatusBadRequest, "invalid sender type")
		return
	}

	ctx := r.Context()
	msg, err := handler.DialogService.AddMessage(
		ctx,
		models2.DialogID(body.DialogID),
		sender,
		body.Content,
	)

	if sender == models2.SenderUser {
		reply := makeAssistantReply(body.Content)

		_, aErr := handler.DialogService.AddMessage(
			ctx,
			models2.DialogID(body.DialogID),
			models2.SenderAssistant,
			reply,
		)
		if aErr != nil {
			log.Println("add assistant message error: ", aErr)
		}
	}

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

	_, err := authhttp.UserIDFromRequest(r)
	if err != nil {
		if errors.Is(err, authhttp.ErrNotAuthHeader) || errors.Is(err, authhttp.ErrInvalidToken) {
			httpresp.Error(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		log.Println("auth error: ", err)
		httpresp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	dialogID := r.URL.Query().Get("dialog_id")
	if dialogID == "" {
		httpresp.Error(w, http.StatusBadRequest, "dialog_id is required")
		return
	}

	ctx := r.Context()
	msg, err := handler.DialogService.ListMessages(ctx, models2.DialogID(dialogID))
	if err != nil {
		log.Println("list messages error: ", err)
		httpresp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	httpresp.JSON(w, http.StatusOK, msg)
}

func makeAssistantReply(userText string) string {
	text := strings.TrimSpace(userText)
	if text == "" {
		return "Я вас слушаю 🙂"
	}

	short := text
	r := []rune(short)
	if len(r) > 180 {
		short = string(r[:180]) + "..."
	}

	if strings.Contains(short, "?") {
		return "Понял вопрос? Коротко: " + short
	}
	return "Принято. Я понял: " + short
}
