package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 5

// Connection
func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

// AllMovies returns all movies from the database
func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {

	// context to handle db query timeout
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// query the database
	query := `
		select
			id, title, release_date, runtime,
			mpaa_rating, description, coalesce(image, ''),
			created_at, updated_at
		from
			movies
		order by
			title
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// declare a slice to hold the models.Movies objects
	var movies []*models.Movie

	// loop through the rows and scan the data into the models.Movies object
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}

// สร้าง Function ค้นหา User จาก Email
func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {

	// context to handle db query timeout
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password,
			created_at, updated_at from users where email = $1`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
