package ports

import (
	"github.com/afikrim/movie-api-go/internal/core/domain"
)

type GenreService interface {
	Create(name string) (domain.Genre, error)
	FindAll() ([]domain.Genre, error)
	FindOne(id int64) (domain.Genre, error)
	Update(id int64, name string) (domain.Genre, error)
	Remove(id int64) (domain.Genre, error)
}
