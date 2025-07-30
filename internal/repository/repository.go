package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	// Get all movies
	AllMovies() ([]*models.Movie, error)
}
