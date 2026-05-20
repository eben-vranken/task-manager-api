-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL
);

-- +goose Down
DROP TABLE users;