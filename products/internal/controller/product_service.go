package controller

import (
	"context"

	"github.com/Binit-Dhakal/e-fit/products/internal/model"
	"github.com/Binit-Dhakal/e-fit/products/internal/ports"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(r ports.ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (p *ProductService) CreateProduct(ctx context.Context, product *model.Product) (string, error) {
	return p.repo.Save(ctx, product)
}

func (p *ProductService) GetProductBySlug(ctx context.Context, slug string) (*model.Product, error) {
	return p.repo.FindBySlug(ctx, slug)
}
