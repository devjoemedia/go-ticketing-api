package utils

import (
	"encoding/json"
	"net/http"

	"github.com/devjoemedia/chitodopostgress/types"
)

func JSON[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Response encoding failed", http.StatusInternalServerError)
	}
}

func Success[T any](w http.ResponseWriter, status int, data T, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := types.SuccessResponse[T]{
		Data:    data,
		Message: msg,
		Success: true,
		Status:  status,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Response encoding failed", http.StatusInternalServerError)
	}
}

func Error(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := types.ErrorResponse{
		Message: msg,
		Success: false,
		Status:  status,
	}
	json.NewEncoder(w).Encode(resp)
}
