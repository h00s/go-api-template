package sql

const GetMovies = `
	SELECT id, title, release_year, director, rating, created_at
	FROM movies
`
