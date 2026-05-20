-- +goose Up
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    task_status status NOT NULL,
    task_priority priority NOT NULL,
    title VARCHAR(64) NOT NULL,
    description VARCHAR(256),
    due_date DATE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
        REFERENCES users (id)
);

-- +goose Down
DROP TABLE tasks;