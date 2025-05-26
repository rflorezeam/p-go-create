package tests

import (
	"errors"
	"testing"

	"github.com/rflorezeam/libro-create/models"
	"github.com/rflorezeam/libro-create/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mock del repositorio
type MockLibroRepository struct {
	mock.Mock
}

func (m *MockLibroRepository) CrearLibro(libro models.Libro) (*mongo.InsertOneResult, error) {
	args := m.Called(libro)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func TestCrearLibro_Exitoso(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)
	service := services.NewLibroService(mockRepo)
	
	libro := models.Libro{
		Titulo: "Test Libro",
		Autor:  "Test Autor",
	}
	
	expectedResult := &mongo.InsertOneResult{}
	mockRepo.On("CrearLibro", libro).Return(expectedResult, nil)

	// Act
	result, err := service.CrearLibro(libro)

	// Assert
	assert.NoError(t, err, "No debería haber error al crear el libro")
	assert.NotNil(t, result, "El resultado no debería ser nulo")
	assert.Equal(t, expectedResult, result, "El resultado debería ser igual al esperado")
	mockRepo.AssertExpectations(t)
}

func TestCrearLibro_Error(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)
	service := services.NewLibroService(mockRepo)
	
	libro := models.Libro{
		Titulo: "Test Libro",
		Autor:  "Test Autor",
	}
	
	expectedError := errors.New("error al crear libro")
	mockRepo.On("CrearLibro", libro).Return(nil, expectedError)

	// Act
	result, err := service.CrearLibro(libro)

	// Assert
	assert.Error(t, err, "Debería haber un error al crear el libro")
	assert.Equal(t, expectedError, err, "El error debería ser igual al esperado")
	assert.Nil(t, result, "El resultado debería ser nulo")
	mockRepo.AssertExpectations(t)
}

func TestCrearLibro_LibroVacio(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)
	service := services.NewLibroService(mockRepo)
	
	libroVacio := models.Libro{}
	expectedResult := &mongo.InsertOneResult{}
	mockRepo.On("CrearLibro", libroVacio).Return(expectedResult, nil)

	// Act
	result, err := service.CrearLibro(libroVacio)

	// Assert
	assert.NoError(t, err, "No debería haber error al crear un libro vacío")
	assert.NotNil(t, result, "El resultado no debería ser nulo")
	mockRepo.AssertExpectations(t)
} 