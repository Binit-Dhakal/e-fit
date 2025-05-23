package grpc

import (
	"context"
	"fmt"

	"github.com/Binit-Dhakal/e-fit/products/gen"
	"github.com/Binit-Dhakal/e-fit/products/internal/model"
	"github.com/Binit-Dhakal/e-fit/products/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedProductServiceServer
	ctrl ports.ProductService
}

func NewHandler(ctrl ports.ProductService) *Handler {
	return &Handler{
		ctrl: ctrl,
	}
}

func (h *Handler) GetProductBySlug(ctx context.Context, req *gen.GetProductBySlugRequest) (*gen.GetProductBySlugResponse, error) {
	if req == nil || req.Slug == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil or empty slug")
	}

	fmt.Println("slug:  ", req.Slug)
	p, err := h.ctrl.GetProductBySlug(ctx, req.Slug)
	if err != nil {
		return nil, err
	}
	fmt.Println(p)
	genProduct, err := model.ProductToProto(p)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.GetProductBySlugResponse{Product: genProduct}, nil
}

func (h *Handler) CreateProduct(ctx context.Context, req *gen.CreateProductRequest) (*gen.CreateProductResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil product request")
	}

	product, err := model.ProtoToProduct(req.Product)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	_, err = h.ctrl.CreateProduct(ctx, product)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.CreateProductResponse{}, nil
}
