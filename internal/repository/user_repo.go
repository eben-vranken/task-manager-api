package repository

import (
	"context"
	"database/sql"

	"github.com/eben-vranken/task-manager-api/internal/models"
)


type UserRepository struct {
	db *sql.DB
}

func (ur UserRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	err := ur.db.QueryRowContext(ctx, `INSERT INTO users (name)
	VALUES ($1)
	RETURNING id
	`, user.Name).Scan(&user.ID)
	
	return user, err
}

func NewUserRepository(db *sql.DB) *UserRepository {
	t := new(UserRepository)
	t.db = db
	return t
}