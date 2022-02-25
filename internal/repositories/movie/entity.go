package movie_repository

import (
	"time"

	genre_repository "github.com/afikrim/movie-api-go/internal/repositories/genre"
)

type Movie struct {
	ID          int64                    `gorm:"column:id;autoIncrement;primaryKey"`
	Name        string                   `gorm:"column:name"`
	Description string                   `gorm:"column:description"`
	Genres      []genre_repository.Genre `gorm:"many2many:movie_genres"`
	CreatedAt   *time.Time               `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   *time.Time               `gorm:"column:updated_at;autoUpdateTime"`
}
