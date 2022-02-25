package movie_repository

import (
	"context"

	"github.com/afikrim/movie-api-go/internal/core/domain"
	genre_repository "github.com/afikrim/movie-api-go/internal/repositories/genre"
	"gorm.io/gorm"
)

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *movieRepository {
	return &movieRepository{
		db: db,
	}
}

func (movieRepository *movieRepository) Migrate() {
	db := movieRepository.db

	db.AutoMigrate(&Movie{})
}

func (movieRepository *movieRepository) Find(ctx context.Context) ([]*domain.Movie, error) {
	query := movieRepository.db.Model(&Movie{}).WithContext(ctx)

	var moviesModel []*Movie
	if err := query.Preload("Genres").Find(&moviesModel).Error; err != nil {
		return nil, err
	}

	var movies []*domain.Movie
	for _, movieModel := range moviesModel {
		var genres []*domain.Genre
		for _, genreModel := range movieModel.Genres {
			genres = append(genres, &domain.Genre{
				ID:        genreModel.ID,
				Name:      genreModel.Name,
				CreatedAt: genreModel.CreatedAt,
				UpdatedAt: genreModel.UpdatedAt,
			})
		}

		movies = append(movies, &domain.Movie{
			ID:          movieModel.ID,
			Name:        movieModel.Name,
			Description: movieModel.Description,
			Genres:      genres,
			CreatedAt:   movieModel.CreatedAt,
			UpdatedAt:   movieModel.UpdatedAt,
		})
	}

	return movies, nil
}

func (movieRepository *movieRepository) FindOne(ctx context.Context, id int64) (*domain.Movie, error) {
	query := movieRepository.db.Model(&Movie{}).WithContext(ctx)

	var movieModel *Movie
	if err := query.Preload("Genres").Where(&Movie{ID: id}).Limit(1).Find(&movieModel).Error; err != nil {
		return nil, err
	}

	var genres []*domain.Genre
	for _, genreModel := range movieModel.Genres {
		genres = append(genres, &domain.Genre{
			ID:        genreModel.ID,
			Name:      genreModel.Name,
			CreatedAt: genreModel.CreatedAt,
			UpdatedAt: genreModel.UpdatedAt,
		})
	}

	var movie *domain.Movie
	movie = &domain.Movie{
		ID:          movieModel.ID,
		Name:        movieModel.Name,
		Description: movieModel.Description,
		Genres:      genres,
		CreatedAt:   movieModel.CreatedAt,
		UpdatedAt:   movieModel.UpdatedAt,
	}

	return movie, nil
}

func (movieRepository *movieRepository) Insert(ctx context.Context, createMovieDto domain.CreateMovieDto) (*domain.Movie, error) {
	var genresModel []genre_repository.Genre
	if err := movieRepository.db.Model(&genre_repository.Genre{}).WithContext(ctx).Find(&genresModel, createMovieDto.Genres).Error; err != nil {
		return nil, err
	}

	movieModel := &Movie{
		Name:        createMovieDto.Name,
		Description: createMovieDto.Description,
		Genres:      genresModel,
	}
	if err := movieRepository.db.WithContext(ctx).Create(&movieModel).Error; err != nil {
		return nil, err
	}

	var genres []*domain.Genre
	for _, genreModel := range movieModel.Genres {
		genres = append(genres, &domain.Genre{
			ID:        genreModel.ID,
			Name:      genreModel.Name,
			CreatedAt: genreModel.CreatedAt,
			UpdatedAt: genreModel.UpdatedAt,
		})
	}

	var movie *domain.Movie
	movie = &domain.Movie{
		ID:          movieModel.ID,
		Name:        movieModel.Name,
		Description: movieModel.Description,
		Genres:      genres,
		CreatedAt:   movieModel.CreatedAt,
		UpdatedAt:   movieModel.UpdatedAt,
	}

	return movie, nil
}

func (movieRepository *movieRepository) Update(ctx context.Context, id int64, updateMovieDto domain.UpdateMovieDto) (*domain.Movie, error) {
	var movie *domain.Movie
	movie, err := movieRepository.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	var genresModel []genre_repository.Genre
	for _, genre := range movie.Genres {
		genresModel = append(genresModel, genre_repository.Genre{
			ID:        genre.ID,
			Name:      genre.Name,
			CreatedAt: genre.CreatedAt,
			UpdatedAt: genre.UpdatedAt,
		})
	}

	movieModel := &Movie{
		ID:          movie.ID,
		Name:        updateMovieDto.Name,
		Description: updateMovieDto.Description,
		Genres:      genresModel,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
	}

	result := movieRepository.db.WithContext(ctx).Where(&Movie{ID: id}).Updates(&movieModel)
	if err := result.Error; err != nil {
		return nil, err
	}

	movie = &domain.Movie{
		ID:          movieModel.ID,
		Name:        movieModel.Name,
		Description: movieModel.Description,
		Genres:      movie.Genres,
		CreatedAt:   movieModel.CreatedAt,
		UpdatedAt:   movieModel.UpdatedAt,
	}

	return movie, nil
}

func (movieRepository *movieRepository) Delete(ctx context.Context, id int64) (*domain.Movie, error) {
	var movie *domain.Movie
	movie, err := movieRepository.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	result := movieRepository.db.WithContext(ctx).Preload("Genres").Where(&Movie{ID: id}).Delete(&Movie{})
	if err := result.Error; err != nil {
		return nil, err
	}

	return movie, nil
}
