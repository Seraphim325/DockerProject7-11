package handler

import (
	"api_server/internal/domain"
	"errors"
	"net/http"
)

func mapToResponse(err error) (int, string) {
	switch {
	case errors.Is(err, domain.ErrBigNumber):
		return http.StatusBadRequest, "invalid input data"
	case errors.Is(err, domain.ErrDuplicate):
		return http.StatusConflict, "data already exists"
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}

func Error(w http.ResponseWriter, err error) {
	status, message := mapToResponse(err)
	http.Error(w, message, status)
}
