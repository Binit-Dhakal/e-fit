package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/Binit-Dhakal/e-fit/users/internal/model"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func New() (*Repository, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return &Repository{
		DB: conn,
	}, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*model.User, error) {
	query := `SELECT id,username, password, email, email_verified from users where id=$1`

	var user model.User
	err := r.DB.QueryRow(ctx, query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.EmailVerified)

	return &user, err
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id,username, password, email, email_verified from users where email=$1`

	var user model.User
	err := r.DB.QueryRow(ctx, query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.EmailVerified)

	return &user, err
}

func (r *Repository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT into users (username, password, email) values ($1,$2,$3) returning id`

	args := []any{user.Username, user.Password, user.Email}

	err := r.DB.QueryRow(ctx, query, args...).Scan(&user.ID)
	if err != nil {
		// Duplicate email check
		return err
	}

	return nil
}
