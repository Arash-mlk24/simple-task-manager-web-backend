package repository

import (
	"context"
	"database/sql"
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/core/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) (int64, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (repository *userRepository) Create(ctx context.Context, user entity.User) (int64, error) {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	var id int64
	err := repository.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&id)
	return id, err
}

func (repository *userRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE id = $1`
	user := &entity.User{}
	err := repository.db.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	return user, err
}

func (repository *userRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	query := `SELECT id, username, email, password FROM users`

	rows, err := repository.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
