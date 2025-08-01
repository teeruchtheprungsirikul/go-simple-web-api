package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB

	// Get User by Email
	GetUserByEmail(email string) (*models.User, error)
	// Get User by ID
	GetUserByID(id int) (*models.User, error)

	// Get all movies
	AllMovies() ([]*models.Movie, error)
	// Get One Movie by ID
	OneMovie(id int) (*models.Movie, error)
	// Edit One Movie by ID
	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)
}
