package service

import (
	"context"

	"golang_clean_architecture/dto/product"
	"golang_clean_architecture/entity"
)

type ProductService interface {
	FindAll(ctx context.Context) []entity.Product
	FindById(ctx context.Context, id int) (entity.Product, error)
	Create(ctx context.Context, request product.ProductCreateRequest) (entity.Product, error)
}
