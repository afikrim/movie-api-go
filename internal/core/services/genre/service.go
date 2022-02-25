package genre_service

import (
	"context"

	"github.com/afikrim/movie-api-go/internal/core/domain"
	"github.com/afikrim/movie-api-go/internal/core/ports"
)

type service struct {
	genreRepository ports.GenreRepository
}

func NewGenreService(genreRepository ports.GenreRepository) *service {
	genreRepository.Migrate()

	return &service{
		genreRepository: genreRepository,
	}
}

func (service *service) Create(ctx context.Context, createGenreDto domain.CreateGenreDto) (*domain.Genre, error) {
	genre, err := service.genreRepository.Insert(ctx, createGenreDto)
	if err != nil {
		return &domain.Genre{}, err
	}

	return genre, nil
}

func (service *service) FindAll(ctx context.Context) ([]*domain.Genre, error) {
	genres, err := service.genreRepository.Find(ctx)
	if err != nil {
		return []*domain.Genre{}, err
	}

	return genres, nil
}

func (service *service) FindOne(ctx context.Context, id int64) (*domain.Genre, error) {
	genre, err := service.genreRepository.FindOne(ctx, id)
	if err != nil {
		return &domain.Genre{}, err
	}

	return genre, nil
}

func (service *service) Update(ctx context.Context, id int64, updateGenreDto domain.UpdateGenreDto) (*domain.Genre, error) {
	genre, err := service.genreRepository.Update(ctx, id, updateGenreDto)
	if err != nil {
		return &domain.Genre{}, err
	}

	return genre, nil
}

func (service *service) Remove(ctx context.Context, id int64) (*domain.Genre, error) {
	genre, err := service.genreRepository.Delete(ctx, id)
	if err != nil {
		return &domain.Genre{}, err
	}

	return genre, nil
}
