package migrations

const CreateMovies = `
	BEGIN;

	CREATE TABLE Movies (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		release_year INT NOT NULL,
		director VARCHAR(255) NOT NULL,
		rating FLOAT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE UNIQUE INDEX movies_title_release_year_idx ON Movies (title, release_year);

	UPDATE schema_migrations SET version = 1 WHERE version = 0;
	
	COMMIT;
`
