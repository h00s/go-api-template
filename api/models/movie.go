package models

import (
	"encoding/json"
	"time"
)

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ReleaseYear int       `json:"release_year"`
	Director    string    `json:"director"`
	Rating      float64   `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
}

func (l *Movie) Marshal() ([]byte, error) {
	return json.Marshal(l)
}

func (l *Movie) Unmarshal(data []byte) error {
	return json.Unmarshal(data, l)
}
