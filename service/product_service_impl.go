package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"

	"golang_clean_architecture/dto/product"
	"golang_clean_architecture/entity"
	"golang_clean_architecture/exception"
	"golang_clean_architecture/helper"
	"golang_clean_architecture/repository"
)

type ProductServiceImpl struct {
	Repository repository.ProductRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewProductService(repository repository.ProductRepository, db *sql.DB) *ProductServiceImpl {
	return &ProductServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validator.New(),
	}
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []entity.Product {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.Repository.FindAll(ctx, tx)
	return products
}

func (service *ProductServiceImpl) FindById(ctx context.Context, id int) (entity.Product, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product, err := service.Repository.FindById(ctx, tx, id)
	if err != nil {
		return product, exception.NewNotFoundError("Product tidak di temukan")
	}

	return product, nil
}

func (service *ProductServiceImpl) Create(ctx context.Context, request product.ProductCreateRequest) (entity.Product, error) {
	if err := service.Validate.Struct(request); err != nil {
		return entity.Product{}, exception.NewValidationError(err.(validator.ValidationErrors))
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := entity.Product{
		Name:  request.Name,
		Stok:  request.Stok,
		Price: request.Price,
	}

	return service.Repository.Create(ctx, tx, product)
}
