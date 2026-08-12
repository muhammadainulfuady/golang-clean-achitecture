package usecase

import (
	"context"
	"golang_clean_architecture/internal/model"
)

type BookUsecase interface {
	Create(ctx context.Context, request model.CreateBookRequest) (model.BookResponse, error)
	Update(ctx context.Context, request model.UpdateBookRequest) (model.BookResponse, error)
	Delete(ctx context.Context, bookId int) error
	FindById(ctx context.Context, bookId int) (model.BookResponse, error)
	FindAll(ctx context.Context) ([]model.BookResponse, error)
}
