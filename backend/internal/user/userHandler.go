package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var userData SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	err = validate.Struct(userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	response, err := h.service.Register(context.Background(), &userData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var userData SignInRequest
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	err = validate.Struct(userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	user, err := h.service.Login(context.Background(), &userData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	authCookie := http.Cookie{
		Name:     "token",
		Value:    user.AccessToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &authCookie)

	response := &SignInResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {

	authCookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &authCookie)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"message": "Logout successfull"})
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Hi %s!", r.Context().Value("user_name"))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}
