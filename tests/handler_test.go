package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rflorezeam/libro-create/handlers"
	"github.com/rflorezeam/libro-create/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockLibroService struct {
	mock.Mock
}

func (m *MockLibroService) CrearLibro(libro models.Libro) (*mongo.InsertOneResult, error) {
	args := m.Called(libro)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func TestCrearLibroHandler_Exitoso(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	libro := models.Libro{
		Titulo: "Test Libro",
		Autor:  "Test Autor",
	}

	body, _ := json.Marshal(libro)
	req := httptest.NewRequest(http.MethodPost, "/libros", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	expectedResult := &mongo.InsertOneResult{}
	mockService.On("CrearLibro", libro).Return(expectedResult, nil)

	// Act
	handler.CrearLibro(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestCrearLibroHandler_ErrorJSON(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	// JSON inválido
	invalidJSON := []byte(`{"titulo": "Test`)
	req := httptest.NewRequest(http.MethodPost, "/libros", bytes.NewBuffer(invalidJSON))
	w := httptest.NewRecorder()

	// Act
	handler.CrearLibro(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCrearLibroHandler_ErrorServicio(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	libro := models.Libro{
		Titulo: "Test Libro",
		Autor:  "Test Autor",
	}

	body, _ := json.Marshal(libro)
	req := httptest.NewRequest(http.MethodPost, "/libros", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	mockService.On("CrearLibro", libro).Return(nil, errors.New("error del servicio"))

	// Act
	handler.CrearLibro(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
} 