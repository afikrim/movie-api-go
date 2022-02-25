package domain

import "time"

type Movie struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Genres      []*Genre   `json:"genres"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type CreateMovieDto struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Genres      []*int64 `json:"genres"`
}

type UpdateMovieDto struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Genres      []*int64 `json:"genres"`
}
