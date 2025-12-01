package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tahmazidik/llm-assistant-web-lab/internal/models"
	userssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/users"
	httpresp "github.com/tahmazidik/llm-assistant-web-lab/internal/transport/http/response"
)

const demoToken = "demo-token-123"

// UserHandler отвечает за HTTP-операции, связанные с пользователями
type UserHandler struct {
	UserService userssvc.Service
}

// NewUserHandler создает новый хендлер пользователей
func NewUserHandler(userService userssvc.Service) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// Структура для приема JSON из запроса
type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// Register обрабатывает регистрацию нового пользователя(POST)
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Разрешаем только POST
	if r.Method != http.MethodPost {
		httpresp.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	//Парсим JSON из тела запроса
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpresp.Error(w, http.StatusBadRequest, "invalid json")
		return
	}

	if req.Email == "" {
		httpresp.Error(w, http.StatusBadRequest, "email is required")
		return
	}

	if req.Password == "" {
		httpresp.Error(w, http.StatusBadRequest, "password is required")
		return
	}

	if req.Name == "" {
		httpresp.Error(w, http.StatusBadRequest, "name is required")
		return
	}

	// Вызываем бизнес-логику
	ctx := r.Context()
	user, err := h.UserService.Register(ctx, req.Email, req.Password, req.Name)
	if err != nil {
		switch err {
		case userssvc.ErrEmailAlreadyInUse:
			httpresp.Error(w, http.StatusConflict, "email already in use")
		default:
			log.Println("error registering user:", err)
			httpresp.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	//Отдаем созданного пользователя
	httpresp.JSON(w, http.StatusCreated, user)
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}

// Login обрабатывает POST /users/login
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpresp.Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpresp.Error(w, http.StatusBadRequest, "invalid json")
		return
	}

	if req.Email == "" {
		httpresp.Error(w, http.StatusBadRequest, "email is required")
		return
	}

	if req.Password == "" {
		httpresp.Error(w, http.StatusBadRequest, "password is required")
		return
	}

	ctx := r.Context()
	user, err := h.UserService.Authenticate(ctx, req.Email, req.Password)
	if err != nil {
		switch err {
		case userssvc.ErrInvalidCredentials:
			httpresp.Error(w, http.StatusUnauthorized, "invalid email or password")
		default:
			log.Println("error authenticating user:", err)
			httpresp.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	resp := loginResponse{
		User:  user,
		Token: demoToken,
	}
	httpresp.JSON(w, http.StatusOK, resp)
}
