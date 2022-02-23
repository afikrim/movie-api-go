package genre_service

import (
	"github.com/afikrim/movie-api-go/internal/core/domain"
	"github.com/afikrim/movie-api-go/internal/core/ports"
)

type service struct {
	genreRepository ports.GenreRepository
}

func New(genreRepository ports.GenreRepository) *service {
	return &service{
		genreRepository: genreRepository,
	}
}

func (service *service) Create(name string) (domain.Genre, error) {
	genre, err := service.genreRepository.Insert(name)
	if err != nil {
		return domain.Genre{}, err
	}

	return genre, nil
}

func (service *service) FindAll() ([]domain.Genre, error) {
	genres, err := service.genreRepository.Find()
	if err != nil {
		return []domain.Genre{}, err
	}

	return genres, nil
}

func (service *service) FindOne(id int64) (domain.Genre, error) {
	genre, err := service.genreRepository.FindOne(id)
	if err != nil {
		return domain.Genre{}, err
	}

	return genre, nil
}

func (service *service) Update(id int64, name string) (domain.Genre, error) {
	genre, err := service.genreRepository.Update(id, name)
	if err != nil {
		return domain.Genre{}, err
	}

	return genre, nil
}

func (service *service) Remove(id int64) (domain.Genre, error) {
	genre, err := service.genreRepository.Delete(id)
	if err != nil {
		return domain.Genre{}, err
	}

	return genre, nil
}
