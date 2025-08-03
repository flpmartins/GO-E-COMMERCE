package handlers

import (
	"encoding/json"
	"net/http"
	"tmp-api/internal/service"
	"tmp-api/pkg/httpx"

	"github.com/go-playground/validator/v10"
)

type CreatePermissionRequest struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type UpdatePermissionRequest struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

type PermissionHandler struct {
	service   service.PermissionService
	validator *validator.Validate
}

func NewPermissionHandler(service service.PermissionService) *PermissionHandler {
	return &PermissionHandler{service: service, validator: validator.New()}
}

func (h *PermissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) {
	var req CreatePermissionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		httpx.WriteValidationErrors(w, err)
		return
	}

	permission, err := h.service.CreatePermission(r.Context(), req.Name, req.Value)
	if err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, permission)
}

func (h *PermissionHandler) GetPermission(w http.ResponseWriter, r *http.Request) {
	id, err := httpx.ParseUUIDParam(r, "id")
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	permission, err := h.service.GetPermissionByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(permission)
}

func (h *PermissionHandler) GetAllPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := h.service.GetAllPermissions(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(permissions) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(permissions)
}

func (h *PermissionHandler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	id, err := httpx.ParseUUIDParam(r, "id")
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	var request UpdatePermissionRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	err = h.service.UpdatePermission(r.Context(), id, request.Name, request.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *PermissionHandler) DeletePermission(w http.ResponseWriter, r *http.Request) {
	id, err := httpx.ParseUUIDParam(r, "id")
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.DeletePermission(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
