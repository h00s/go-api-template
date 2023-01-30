package models

import (
	"context"

	"github.com/h00s/go-api-template/api/services"
	"github.com/h00s/go-api-template/db/sql"
)

type Movies struct {
	services *services.Services
}

func NewMovies(services *services.Services) *Movies {
	return &Movies{
		services: services,
	}
}

func (m *Movies) FindAll() ([]Movie, error) {
	rows, err := m.services.DB.Conn.Query(context.Background(), sql.GetMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.Director, &movie.Rating, &movie.CreatedAt)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}
