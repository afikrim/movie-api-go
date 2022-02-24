package ports

import (
	"context"

	"github.com/afikrim/movie-api-go/internal/core/domain"
)

type GenreRepository interface {
	Migrate()
	Find(ctx context.Context) ([]*domain.Genre, error)
	FindOne(ctx context.Context, id int64) (*domain.Genre, error)
	Insert(ctx context.Context, createGenreDto domain.CreateGenreDto) (*domain.Genre, error)
	Update(ctx context.Context, id int64, updateGenreDto domain.UpdateGenreDto) (*domain.Genre, error)
	Delete(ctx context.Context, id int64) (*domain.Genre, error)
}
