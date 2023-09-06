package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/kensparta/fullcycle-multithreading/dto"
	"github.com/kensparta/fullcycle-multithreading/requests"
	"net/http"
)

func GetCep(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cep := chi.URLParam(r, "cep")
	if !isValidCep(cep) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{
			Message: "cep format error, should be something like: 00000-000",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := requests.SelectResponse(ctx, cep, w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
}
