package handlers

import (
	"encoding/json"
	"net/http"
	"tmp-api/internal/service"
	"tmp-api/pkg/httpx"

	"github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=6"`
	IdPermission string `json:"id_permission" validate:"required,uuid4"`
}
type UpdateUserRequest struct {
	Name         *string `json:"name"`
	Email        *string `json:"email"`
	IdPermission *string `json:"id_permission"`
}
type UserHandler struct {
	service   service.IUserService
	validator *validator.Validate
}

func NewUserHandler(s service.IUserService) *UserHandler {
	return &UserHandler{service: s, validator: validator.New()}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := handler.validator.Struct(req); err != nil {
		httpx.WriteValidationErrors(w, err)
		return
	}

	user, err := handler.service.CreateUser(r.Context(), req.Name, req.Email, req.IdPermission, req.Password)

	if err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, user)
}

func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := httpx.ParseUUIDParam(r, "id")
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := handler.service.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (handler *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handler.service.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var request UpdateUserRequest

	id, err := httpx.ParseUUIDParam(r, "id")
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = handler.service.UpdateUser(r.Context(), id, request.Name, request.Email, request.IdPermission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := httpx.ParseUUIDParam(r, "id")
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = handler.service.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
