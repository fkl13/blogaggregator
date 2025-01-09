-- +goose Up
CREATE TABLE users (
	id uuid ,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	name text NOT NULL,
	PRIMARY KEY (id),
	UNIQUE(name)
);

-- +goose Down
DROP TABLE users;
