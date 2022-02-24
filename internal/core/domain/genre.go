package domain

import "time"

type Genre struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CreateGenreDto struct {
	Name string `json:"name"`
}

type UpdateGenreDto struct {
	Name string `json:"name"`
}
