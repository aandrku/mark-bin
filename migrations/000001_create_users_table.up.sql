CREATE TABLE IF NOT EXISTS users (
	id BIGSERIAL PRIMARY KEY,
	username VARCHAR(32),
	email VARCHAR(255),
	hashed_password CHAR(60),
	created TIMESTAMP
);
