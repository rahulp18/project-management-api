package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) CreateUser(ctx context.Context, user UserInput) (User, error) {
	query := `INSERT INTO users (name,email,password)
          VALUES ($1,$2,$3)
		  RETURNING id,name,email,created_at,updated_at
  `
	var data User
	err := repo.db.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&data.ID, &data.Name, &data.Email, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return User{}, err
	}
	return data, nil

}
func (repo *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}

	query := `
		SELECT id, email, name, password, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	err := repo.db.QueryRow(ctx, query, email).
		Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
