package handlers

import (
	error2 "Auth-Service/internal/domain"
	"Auth-Service/internal/dtos"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

func HandleHttpError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, error2.ErrInvalidInput):
		respondError(w, http.StatusBadRequest, err.Error())

	case errors.Is(err, error2.ErrAlreadyExists):
		respondError(w, http.StatusConflict, err.Error())

	case errors.Is(err, error2.ErrUnauthorized):
		respondError(w, http.StatusUnauthorized, err.Error())

	case errors.Is(err, error2.ErrForbidden):
		respondError(w, http.StatusForbidden, err.Error())

	case errors.Is(err, error2.ErrNotFound):
		respondError(w, http.StatusNotFound, err.Error())

	default:
		log.Println("internal error:", err)
		respondError(w, http.StatusInternalServerError, "internal server error")
	}
}

func respondError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(dtos.Error{
		Code:        strconv.Itoa(code),
		Description: msg,
	})
}
