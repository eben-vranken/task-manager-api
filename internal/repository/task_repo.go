package repository

import (
	"context"
	"database/sql"

	"github.com/eben-vranken/task-manager-api/internal/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func (tr TaskRepository) Create(ctx context.Context, task models.Task) (models.Task, error) {
	execErr := tr.DB.QueryRowContext(ctx, `INSERT INTO tasks 
	(user_id,
	 task_status,
	 task_priority, 
	 title, 
	 description, 
	 due_date) 
	 VALUES ($1, $2, $3, $4, $5, $6)
	 RETURNING id, created_at, updated_at`,
		task.UserID, task.TaskStatus, task.TaskPriority, task.Title, task.Description, task.DueDate,
	).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)

	return task, execErr
}
