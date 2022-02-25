package ports

import (
	"context"

	"github.com/afikrim/movie-api-go/internal/core/domain"
)

type GenreService interface {
	Create(ctx context.Context, createGenreDto domain.CreateGenreDto) (*domain.Genre, error)
	FindAll(ctx context.Context) ([]*domain.Genre, error)
	FindOne(ctx context.Context, id int64) (*domain.Genre, error)
	Update(ctx context.Context, id int64, updateGenreDto domain.UpdateGenreDto) (*domain.Genre, error)
	Remove(ctx context.Context, id int64) (*domain.Genre, error)
}

type MovieService interface {
	Create(ctx context.Context, createMovieDto domain.CreateMovieDto) (*domain.Movie, error)
	FindAll(ctx context.Context) ([]*domain.Movie, error)
	FindOne(ctx context.Context, id int64) (*domain.Movie, error)
	Update(ctx context.Context, id int64, updateMovieDto domain.UpdateMovieDto) (*domain.Movie, error)
	Remove(ctx context.Context, id int64) (*domain.Movie, error)
}
