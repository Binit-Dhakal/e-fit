package model

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type Product struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	Category    string         `json:"category"`
	Images      []string       `json:"images"`
	Brand       string         `json:"brand"`
	Description string         `json:"description"`
	Stock       int            `json:"stock"`
	Price       pgtype.Numeric `json:"price"`
	Rating      pgtype.Numeric `json:"rating"`
	NumReviews  int            `json:"num_reviews"`
	IsFeatured  bool           `json:"is_featured"`
	CreatedAt   time.Time      `json:"-" db:"created_at"`
	UpdatedAt   time.Time      `json:"-" db:"updated_at"`
}
