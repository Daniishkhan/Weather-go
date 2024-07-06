package repository

import (
	"context"
	"database/sql"
	"goweather/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
        INSERT INTO users (username, password_hash, default_location, preferred_unit)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at`

	return r.db.QueryRowContext(ctx, query,
		user.Username, user.PasswordHash, user.DefaultLocation, user.PreferredUnit,
	).Scan(&user.ID, &user.CreatedAt)
}

// Add more methods like GetUserByID, UpdateUser, DeleteUser, etc.
