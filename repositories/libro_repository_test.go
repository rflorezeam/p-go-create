package repositories

import (
	"context"
	"errors"
	"testing"

	"github.com/rflorezeam/libro-create/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

// MockCollection es un mock para la colección de MongoDB
type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

type mockLibroRepository struct {
	collection *MockCollection
}

func NewMockLibroRepository(collection *MockCollection) LibroRepository {
	return &mockLibroRepository{
		collection: collection,
	}
}

func (r *mockLibroRepository) CrearLibro(libro models.Libro) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.TODO(), libro)
}

func TestCrearLibro_Exitoso(t *testing.T) {
	// Arrange
	mockCollection := new(MockCollection)
	repo := NewMockLibroRepository(mockCollection)
	
	libro := models.Libro{
		Titulo: "Test Libro",
		Autor:  "Test Autor",
	}
	
	expectedResult := &mongo.InsertOneResult{}
	mockCollection.On("InsertOne", context.TODO(), libro).Return(expectedResult, nil)

	// Act
	result, err := repo.CrearLibro(libro)

	// Assert
	assert.NoError(t, err, "No debería haber error al crear el libro")
	assert.Equal(t, expectedResult, result, "El resultado debería ser igual al esperado")
	mockCollection.AssertExpectations(t)
}

func TestCrearLibro_ErrorDB(t *testing.T) {
	// Arrange
	mockCollection := new(MockCollection)
	repo := NewMockLibroRepository(mockCollection)
	
	libro := models.Libro{
		Titulo: "Test Libro",
		Autor:  "Test Autor",
	}
	
	expectedError := errors.New("error de conexión con la base de datos")
	mockCollection.On("InsertOne", context.TODO(), libro).Return(nil, expectedError)

	// Act
	result, err := repo.CrearLibro(libro)

	// Assert
	assert.Error(t, err, "Debería haber un error al crear el libro")
	assert.Equal(t, expectedError, err, "El error debería ser igual al esperado")
	assert.Nil(t, result, "El resultado debería ser nulo")
	mockCollection.AssertExpectations(t)
}

func TestCrearLibro_LibroVacio(t *testing.T) {
	// Arrange
	mockCollection := new(MockCollection)
	repo := NewMockLibroRepository(mockCollection)
	
	libroVacio := models.Libro{}
	expectedResult := &mongo.InsertOneResult{}
	mockCollection.On("InsertOne", context.TODO(), libroVacio).Return(expectedResult, nil)

	// Act
	result, err := repo.CrearLibro(libroVacio)

	// Assert
	assert.NoError(t, err, "No debería haber error al crear un libro vacío")
	assert.NotNil(t, result, "El resultado no debería ser nulo")
	mockCollection.AssertExpectations(t)
} 