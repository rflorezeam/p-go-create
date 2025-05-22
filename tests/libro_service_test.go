package tests

import (
	"testing"

	"github.com/rflorezeam/libro-create/models"
	"github.com/rflorezeam/libro-create/repositories"
	"github.com/rflorezeam/libro-create/services"
)

// Mock del repositorio
type mockLibroRepository struct {
	repositories.LibroRepository
}

func (m *mockLibroRepository) CrearLibro(libro models.Libro) (*mongo.InsertOneResult, error) {
	// Simular una inserción exitosa
	return &mongo.InsertOneResult{}, nil
}

func TestCrearLibro(t *testing.T) {
	// Arrange
	mockRepo := &mockLibroRepository{}
	service := services.NewLibroService(mockRepo)
	
	libro := models.Libro{
		Titulo: "Test Libro",
		Autor:  "Test Autor",
	}

	// Act
	result, err := service.CrearLibro(libro)

	// Assert
	if err != nil {
		t.Errorf("Se esperaba nil error, se obtuvo %v", err)
	}
	if result == nil {
		t.Error("Se esperaba un resultado no nulo")
	}
} 