package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/Binit-Dhakal/e-fit/products/internal/model"
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

func (r *Repository) Save(ctx context.Context, product *model.Product) (string, error) {
	query := `
		INSERT INTO products (name, slug, category,images,brand,description, stock,price,is_featured) 
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) 
		returning id `
	args := []any{
		product.Name,
		product.Slug,
		product.Category,
		product.Images,
		product.Brand,
		product.Description,
		product.Stock,
		product.Price,
		product.IsFeatured,
	}
	err := r.DB.QueryRow(ctx, query, args...).Scan(&product.Id)
	return product.Id, err
}

func (r *Repository) FindBySlug(ctx context.Context, slug string) (*model.Product, error) {
	query := `SELECT * FROM products WHERE slug=$1`

	rows, _ := r.DB.Query(ctx, query, slug)
	product, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[model.Product])
	return &product, err
}
