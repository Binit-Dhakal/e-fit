package model

import (
	"github.com/Binit-Dhakal/e-fit/products/gen"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProductToProto(p *Product) (*gen.Product, error) {
	protoPrice, err := p.Price.Value()
	if err != nil {
		return nil, err
	}

	protoRating, err := p.Rating.Value()
	if err != nil {
		return nil, err
	}
	return &gen.Product{
		Id:          p.Id,
		Name:        p.Name,
		Slug:        p.Slug,
		Category:    p.Category,
		Images:      p.Images,
		Brand:       p.Brand,
		Description: p.Description,
		Stock:       int32(p.Stock),
		Price:       protoPrice.(string),
		Rating:      protoRating.(string),
		NumReviews:  int32(p.NumReviews),
		IsFeatured:  p.IsFeatured,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		UpdatedAt:   timestamppb.New(p.UpdatedAt),
	}, nil
}

func ProtoToProduct(p *gen.Product) (*Product, error) {
	var rating pgtype.Numeric
	err := rating.ScanScientific(p.Rating)
	if err != nil {
		return nil, err
	}

	var price pgtype.Numeric
	err = price.ScanScientific(p.Price)
	if err != nil {
		return nil, err
	}

	return &Product{
		Id:          p.Id,
		Name:        p.Name,
		Slug:        p.Slug,
		Category:    p.Category,
		Images:      p.Images,
		Brand:       p.Brand,
		Description: p.Description,
		Stock:       int(p.Stock),
		Price:       price,
		Rating:      rating,
		NumReviews:  int(p.NumReviews),
		IsFeatured:  p.IsFeatured,
		CreatedAt:   p.CreatedAt.AsTime(),
		UpdatedAt:   p.UpdatedAt.AsTime(),
	}, nil
}
