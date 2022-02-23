package ports

import (
	"github.com/afikrim/movie-api-go/internal/core/domain"
)

type GenreRepository interface {
	Find() ([]domain.Genre, error)
	FindOne(id int64) (domain.Genre, error)
	Insert(name string) (domain.Genre, error)
	Update(id int64, name string) (domain.Genre, error)
	Delete(id int64) (domain.Genre, error)
}
