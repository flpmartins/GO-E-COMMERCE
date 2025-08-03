package httpx

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var (
	ErrParamNotFound = errors.New("param not found")
	ErrInvalidUUID   = errors.New("invalid UUID format")
)

func ParseUUIDParam(r *http.Request, param string) (uuid.UUID, error) {
	idStr := chi.URLParam(r, param)
	if idStr == "" {
		return uuid.Nil, ErrParamNotFound
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, ErrInvalidUUID
	}

	return id, nil
}
