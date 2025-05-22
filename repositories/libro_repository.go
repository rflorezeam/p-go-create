package repositories

import (
	"context"

	"github.com/rflorezeam/libro-create/config"
	"github.com/rflorezeam/libro-create/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibroRepository interface {
	CrearLibro(libro models.Libro) (*mongo.InsertOneResult, error)
}

type libroRepository struct {
	collection *mongo.Collection
}

func NewLibroRepository() LibroRepository {
	return &libroRepository{
		collection: config.GetCollection(),
	}
}

func (r *libroRepository) CrearLibro(libro models.Libro) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.TODO(), libro)
} 