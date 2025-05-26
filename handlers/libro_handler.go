package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rflorezeam/libro-create/models"
	"github.com/rflorezeam/libro-create/services"
)

type Handler struct {
	service services.LibroService
}

func NewHandler(service services.LibroService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CrearLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var libro models.Libro
	err := json.NewDecoder(r.Body).Decode(&libro)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error al decodificar el JSON"})
		return
	}
	
	result, err := h.service.CrearLibro(libro)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
} 