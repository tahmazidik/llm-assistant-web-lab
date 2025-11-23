package users

import (
	"encoding/json"
	"log"
	"net/http"

	userssvc "github.com/tahmazidik/llm-assistant-web-lab/internal/services/users"
)

// UserHandler отвечает за HTTP-операции, свяанные с пользователями
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
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//Парсим JSON из тела запроса
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// Вызываем бизнес-логику
	ctx := r.Context()
	user, err := h.UserService.Register(ctx, req.Email, req.Password, req.Name)
	if err != nil {
		switch err {
		case userssvc.ErrEmailAlreadyInUse:
			http.Error(w, "email already in use", http.StatusConflict)
		default:
			log.Println("error registering user:", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	//Отдаем созданного пользователя
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("write response error: ", err)
	}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login обрабатывает POST /users/login
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	user, err := h.UserService.Authenticate(ctx, req.Email, req.Password)
	if err != nil {
		switch err {
		case userssvc.ErrInvalidCredentials:
			http.Error(w, "invalid email or password", http.StatusUnauthorized)
		default:
			log.Println("error authenticating user:", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("write response error: ", err)
	}
}
