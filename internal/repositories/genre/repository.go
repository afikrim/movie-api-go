package genre_repository

import (
	"context"

	"github.com/afikrim/movie-api-go/internal/core/domain"
	"gorm.io/gorm"
)

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) *genreRepository {
	return &genreRepository{
		db: db,
	}
}

func (genreRepository *genreRepository) Migrate() {
	db := genreRepository.db

	db.AutoMigrate(&Genre{})
}

func (genreRepository *genreRepository) Find(ctx context.Context) ([]*domain.Genre, error) {
	query := genreRepository.db.Model(&Genre{}).WithContext(ctx)

	var genresModel []*Genre
	if err := query.Find(&genresModel).Error; err != nil {
		return nil, err
	}

	var genres []*domain.Genre
	for _, genreModel := range genresModel {
		genres = append(genres, &domain.Genre{
			ID:        genreModel.ID,
			Name:      genreModel.Name,
			CreatedAt: genreModel.CreatedAt,
			UpdatedAt: genreModel.UpdatedAt,
		})
	}

	return genres, nil
}

func (genreRepository *genreRepository) FindOne(ctx context.Context, id int64) (*domain.Genre, error) {
	query := genreRepository.db.Model(&Genre{}).WithContext(ctx)

	var genreModel *Genre
	if err := query.Where(&Genre{ID: id}).Limit(1).Find(&genreModel).Error; err != nil {
		return nil, err
	}

	var genre *domain.Genre
	genre = &domain.Genre{
		ID:        genreModel.ID,
		Name:      genreModel.Name,
		CreatedAt: genreModel.CreatedAt,
		UpdatedAt: genreModel.UpdatedAt,
	}

	return genre, nil
}

func (genreRepository *genreRepository) Insert(ctx context.Context, createGenreDto domain.CreateGenreDto) (*domain.Genre, error) {
	genreModel := &Genre{
		Name: createGenreDto.Name,
	}

	result := genreRepository.db.WithContext(ctx).Create(&genreModel)
	if err := result.Error; err != nil {
		return nil, err
	}

	var genre *domain.Genre
	genre = &domain.Genre{
		ID:        genreModel.ID,
		Name:      genreModel.Name,
		CreatedAt: genreModel.CreatedAt,
		UpdatedAt: genreModel.UpdatedAt,
	}

	return genre, nil
}

func (genreRepository *genreRepository) Update(ctx context.Context, id int64, updateGenreDto domain.UpdateGenreDto) (*domain.Genre, error) {
	var genre *domain.Genre
	genre, err := genreRepository.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	genreModel := &Genre{
		ID:        genre.ID,
		Name:      updateGenreDto.Name,
		CreatedAt: genre.CreatedAt,
		UpdatedAt: genre.UpdatedAt,
	}

	result := genreRepository.db.WithContext(ctx).Where(&Genre{ID: id}).Updates(&genreModel)
	if err := result.Error; err != nil {
		return nil, err
	}

	return genre, nil
}

func (genreRepository *genreRepository) Delete(ctx context.Context, id int64) (*domain.Genre, error) {
	var genre *domain.Genre
	genre, err := genreRepository.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	result := genreRepository.db.WithContext(ctx).Where(&Genre{ID: id}).Delete(&Genre{})
	if err := result.Error; err != nil {
		return nil, err
	}

	return genre, nil
}
