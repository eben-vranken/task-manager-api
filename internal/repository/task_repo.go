package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eben-vranken/task-manager-api/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func (tr TaskRepository) Create(ctx context.Context, task models.Task) (models.Task, error) {
	execErr := tr.db.QueryRowContext(ctx, `INSERT INTO tasks 
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

func (tr TaskRepository) GetAll(ctx context.Context) ([]models.Task, error) {
	rows, err := tr.db.QueryContext(ctx, `SELECT 
	id, 
	user_id, 
	task_status, 
	task_priority, 
	title, 
	description, 
	due_date, 
	created_at, 
	updated_at
	FROM tasks;`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.UserID, &task.TaskStatus, &task.TaskPriority, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		} else {
			tasks = append(tasks, task)
		}
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	fmt.Println(tasks)

	return tasks, rows.Err()
}

func (tr TaskRepository) GetSpecificById(ctx context.Context, id string) (*models.Task, error) {
	row, err := tr.db.QueryContext(ctx, `SELECT 
	id, 
	user_id, 
	task_status, 
	task_priority, 
	title, 
	description, 
	due_date, 
	created_at, 
	updated_at
	FROM tasks
	WHERE id = $1;`, id)

	if err != nil {
		return nil, err
	}

	var task models.Task

	for row.Next() {
		err := row.Scan(&task.ID, &task.UserID, &task.TaskStatus, &task.TaskPriority, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	if row.Err() != nil {
		return nil, row.Err()
	}

	return &task, row.Err()
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	t := new(TaskRepository)
	t.db = db
	return t
}
