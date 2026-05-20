-- +goose Up
CREATE TYPE status AS ENUM ('to-do', 'doing', 'done');

-- +goose Down
DROP TYPE status;