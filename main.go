package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rflorezeam/libro-create/config"
	"github.com/rflorezeam/libro-create/models"
	"github.com/rflorezeam/libro-create/repositories"
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
	_ = json.NewDecoder(r.Body).Decode(&libro)
	
	result, err := h.service.CrearLibro(libro)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	// Inicializar la base de datos
	config.ConectarDB()

	// Inicializar las capas
	repo := repositories.NewLibroRepository()
	service := services.NewLibroService(repo)
	handler := NewHandler(service)
	
	// Configurar el router
	router := mux.NewRouter()
	router.HandleFunc("/libros", handler.CrearLibro).Methods("POST")

	// Configurar el puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	fmt.Printf("Servicio de creación de libros corriendo en puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
} 