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

type MovieRepository interface {
	Migrate()
	Find(ctx context.Context) ([]*domain.Movie, error)
	FindOne(ctx context.Context, id int64) (*domain.Movie, error)
	Insert(ctx context.Context, createMovieDto domain.CreateMovieDto) (*domain.Movie, error)
	Update(ctx context.Context, id int64, updateMovieDto domain.UpdateMovieDto) (*domain.Movie, error)
	Delete(ctx context.Context, id int64) (*domain.Movie, error)
}
