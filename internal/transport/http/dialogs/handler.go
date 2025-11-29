package dialogs

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/models"
	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/dialogs"
	httpresp "github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http/response"
)

// временный демо-пользователь, пока нет реальной авторизации
const demoUserID = models.UserID("demo-user-1")

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
	Title string `json:"title"`
}

// Create обрабатывает POST/dialogs запрос для создания нового диалога
func (handler *Handler) Create(w http.ResponseWriter, r *http.Request) {
	// Разрешаем только POST
	if r.Method != http.MethodPost {
		httpresp.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Парсим JSON из тела запроса
	var body createdDialogResponse
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpresp.Error(w, http.StatusBadRequest, "invalid json")
		return
	}

	if body.Title == "" {
		httpresp.Error(w, http.StatusBadRequest, "title is required")
		return
	}

	// Вызываем бизнес-логику для создания диалога
	ctx := r.Context()
	dialog, err := handler.DialogService.CreateDialog(
		ctx,
		demoUserID,
		body.Title,
	)

	if err != nil {
		switch err {
		case dialogssvc.ErrEmptyDialogTitle:
			httpresp.Error(w, http.StatusBadRequest, "title cannot be empty")
		default:
			log.Println("create dialog error: ", err)
			httpresp.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	httpresp.JSON(w, http.StatusCreated, dialog)
}

// List обрабатывает GET/dialogs/list запрос для получения списка диалогов пользователя
func (handler *Handler) List(w http.ResponseWriter, r *http.Request) {
	// Разрешаем только GET
	if r.Method != http.MethodGet {
		httpresp.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	ctx := r.Context()
	dialogs, err := handler.DialogService.ListDialogs(ctx, demoUserID)
	if err != nil {
		log.Println("list dialogs error: ", err)
		httpresp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	httpresp.JSON(w, http.StatusOK, dialogs)
}
