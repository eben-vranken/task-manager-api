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

func (ur UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	rows, err := ur.db.QueryContext(ctx, `SELECT id, name FROM users`)

	if err != nil {
		return nil, err
	}

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, rows.Err()
}

func (ur UserRepository) GetSpecificById(ctx context.Context, id string) (*models.User, error) {

	row, err := ur.db.QueryContext(ctx, "SELECT id, name FROM users WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	var user models.User
	for row.Next() {
		err := row.Scan(&user.ID, &user.Name)

		if err != nil {
			return nil, err
		}
	}

	if row.Err() != nil {
		return nil, err
	}

	return &user, err
}

func NewUserRepository(db *sql.DB) *UserRepository {
	t := new(UserRepository)
	t.db = db
	return t
}
