-- +goose Up
CREATE TABLE feeds (
	id uuid PRIMARY KEY,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	name text NOT NULL,
	url text NOT NULL UNIQUE,
	user_id uuid NOT NULL,
	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
