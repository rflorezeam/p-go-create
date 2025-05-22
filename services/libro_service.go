package services

import (
	"github.com/rflorezeam/libro-create/models"
	"github.com/rflorezeam/libro-create/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibroService interface {
	CrearLibro(libro models.Libro) (*mongo.InsertOneResult, error)
}

type libroService struct {
	repo repositories.LibroRepository
}

func NewLibroService(repo repositories.LibroRepository) LibroService {
	return &libroService{
		repo: repo,
	}
}

func (s *libroService) CrearLibro(libro models.Libro) (*mongo.InsertOneResult, error) {
	return s.repo.CrearLibro(libro)
} 