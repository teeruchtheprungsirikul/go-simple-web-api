package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	// Get all movies
	AllMovies() ([]*models.Movie, error)
	// Get User by Email
	GetUserByEmail(email string) (*models.User, error)
	// Get User by ID
	GetUserByID(id int) (*models.User, error)
}
