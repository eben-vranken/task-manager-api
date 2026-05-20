package models

import "time"

type Task struct {
	ID           int        `json:"id"`
	UserID       int        `json:"user_id"`
	TaskStatus   Status     `json:"task_status"`
	TaskPriority Priority   `json:"task_priority"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	DueDate      *time.Time `json:"due_date"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
