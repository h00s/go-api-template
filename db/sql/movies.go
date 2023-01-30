package sql

const GetMovies = `
	SELECT id, title, release_year, director, rating, created_at
	FROM movies
`

const CreateMovie = `
	INSERT INTO movies (title, release_year, director, rating)
	VALUES ($1, $2, $3, $4)
`
