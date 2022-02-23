package genre_repository

import (
	"database/sql"
	"errors"

	"github.com/afikrim/movie-api-go/internal/core/domain"
)

type genreRepository struct {
	Connection *sql.DB
}

func NewGenreRepository(connection *sql.DB) *genreRepository {
	return &genreRepository{
		Connection: connection,
	}
}

func (genreRepository *genreRepository) Find() ([]domain.Genre, error) {
	query := "SELECT * FROM genres ORDER BY id DESC"
	rows, err := genreRepository.Connection.Query(query)
	if err != nil {
		return []domain.Genre{}, err
	}

	var genres []domain.Genre
	for rows.Next() {
		var genre domain.Genre
		err = rows.Scan(&genre.ID, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt)

		genres = append(genres, genre)
	}

	return genres, err
}

func (genreRepository *genreRepository) FindOne(id int64) (domain.Genre, error) {
	query := "SELECT COUNT(*) as count FROM genres WHERE id = ? LIMIT 1"
	rows, err := genreRepository.Connection.Query(query, id)
	if err != nil {
		return domain.Genre{}, err
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
	}
	if count < 1 {
		return domain.Genre{}, err
	}

	query = "SELECT * FROM genres WHERE id = ? LIMIT 1"
	rows, err = genreRepository.Connection.Query(query, id)
	if err != nil {
		return domain.Genre{}, err
	}

	var genre domain.Genre
	for rows.Next() {
		err = rows.Scan(&genre.ID, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt)
	}
	if err != nil {
		return domain.Genre{}, err
	}

	return genre, err
}

func (genreRepository *genreRepository) Insert(name string) (domain.Genre, error) {
	query := "INSERT INTO genres (name) VALUES (?)"
	statement, err := genreRepository.Connection.Prepare(query)
	if err != nil {
		return domain.Genre{}, err
	}
	defer statement.Close()

	res, err := statement.Exec(name)
	if err != nil {
		return domain.Genre{}, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return domain.Genre{}, err
	}

	genre, err := genreRepository.FindOne(lastID)
	if err != nil {
		return domain.Genre{}, err
	}

	return genre, nil
}

func (genreRepository *genreRepository) Update(id int64, name string) (domain.Genre, error) {
	query := "UPDATE genres SET name = ? WHERE id = ?"
	statement, err := genreRepository.Connection.Prepare(query)
	if err != nil {
		return domain.Genre{}, err
	}
	defer statement.Close()

	res, err := statement.Exec(name, id)
	if err != nil {
		return domain.Genre{}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return domain.Genre{}, err
	}
	if rowsAffected < 1 {
		return domain.Genre{}, errors.New("Failed to update.")
	}

	genre, err := genreRepository.FindOne(id)
	if err != nil {
		return domain.Genre{}, err
	}

	return genre, nil
}

func (genreRepository *genreRepository) Delete(id int64) (domain.Genre, error) {
	genre, err := genreRepository.FindOne(id)
	if err != nil {
		return domain.Genre{}, err
	}

	query := "DELETE FROM genres WHERE id = ?"
	statement, err := genreRepository.Connection.Prepare(query)
	if err != nil {
		return domain.Genre{}, err
	}
	defer statement.Close()

	res, err := statement.Exec(id)
	if err != nil {
		return domain.Genre{}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return domain.Genre{}, err
	}
	if rowsAffected < 1 {
		return domain.Genre{}, errors.New("Failed to delete.")
	}

	return genre, err
}
