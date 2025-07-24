CREATE TABLE IF NOT EXISTS snippets (
	id BIGSERIAL PRIMARY KEY,
	title VARCHAR(64),
	content TEXT,
	created TIMESTAMP,
	updated TIMESTAMP,
	user_id BIGSERIAL
);

ALTER TABLE snippets 
ADD CONSTRAINT fk_snip_user_id
FOREIGN KEY (user_id) REFERENCES users(id);

