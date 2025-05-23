package ports

import (
	"context"

	"github.com/Binit-Dhakal/e-fit/products/internal/model"
)

// Application Port
type ProductService interface {
	CreateProduct(ctx context.Context, p *model.Product) (string, error)
	GetProductBySlug(ctx context.Context, slug string) (*model.Product, error)
}

// Domain Persistance Port
type ProductRepository interface {
	Save(ctx context.Context, p *model.Product) (string, error)
	FindBySlug(ctx context.Context, id string) (*model.Product, error)
}
