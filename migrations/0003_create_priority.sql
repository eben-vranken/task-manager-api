-- +goose Up
CREATE TYPE priority AS ENUM ('low', 'medium', 'high', 'urgent');

-- +goose Down
DROP TYPE priority;