package handlers

import (
	"encoding/json"
	"net/http"
	"tmp-api/internal/service"
	"tmp-api/pkg/httpx"

	"github.com/go-playground/validator/v10"
)

type AuthUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthHandler struct {
	service   service.IAuthService
	validator *validator.Validate
}

func NewAuthHandler(service service.IAuthService) *AuthHandler {
	return &AuthHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *AuthHandler) AuthUser(w http.ResponseWriter, r *http.Request) {
	var req AuthUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		httpx.WriteValidationErrors(w, err)
		return
	}

	user, err := h.service.AuthLogin(r.Context(), req.Email, req.Password)

	if err != nil {
		httpx.WriteError(w, http.StatusUnauthorized, "E-mail ou senha inválidos")
		return
	}

	httpx.WriteJSON(w, http.StatusOK, user)
}
