package repository

import "database/sql"

type TaskRepository struct {
	db *sql.DB
}

func (tr TaskRepository) Create() (db, err) {
	return db, err
}
