package movie_service

import (
	"context"

	"github.com/afikrim/movie-api-go/internal/core/domain"
	"github.com/afikrim/movie-api-go/internal/core/ports"
)

type service struct {
	movieRepository ports.MovieRepository
}

func NewMovieService(movieRepository ports.MovieRepository) *service {
	movieRepository.Migrate()

	return &service{
		movieRepository: movieRepository,
	}
}

func (movieService *service) Create(ctx context.Context, createMovieDto domain.CreateMovieDto) (*domain.Movie, error) {
	movie, err := movieService.movieRepository.Insert(ctx, createMovieDto)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (movieService *service) FindAll(ctx context.Context) ([]*domain.Movie, error) {
	movies, err := movieService.movieRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (movieService *service) FindOne(ctx context.Context, id int64) (*domain.Movie, error) {
	movie, err := movieService.movieRepository.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (movieService *service) Update(ctx context.Context, id int64, updateMovieDto domain.UpdateMovieDto) (*domain.Movie, error) {
	movie, err := movieService.movieRepository.Update(ctx, id, updateMovieDto)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (movieService *service) Remove(ctx context.Context, id int64) (*domain.Movie, error) {
	movie, err := movieService.movieRepository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return movie, nil
}
