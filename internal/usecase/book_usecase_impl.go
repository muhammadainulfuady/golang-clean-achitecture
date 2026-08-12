package usecase

import (
	"context"
	"database/sql"

	"golang_clean_architecture/internal/entity"
	"golang_clean_architecture/internal/model"
	"golang_clean_architecture/internal/model/converter"
	"golang_clean_architecture/internal/repository"

	"github.com/go-playground/validator/v10"
)

type BookUsecaseImpl struct {
	DB             *sql.DB
	BookRepository repository.BookRepository
	Validate       *validator.Validate
}

func NewBookUsecase(db *sql.DB, bookRepository repository.BookRepository, validate *validator.Validate) BookUsecase {
	return &BookUsecaseImpl{
		DB:             db,
		BookRepository: bookRepository,
		Validate:       validate,
	}
}

func (usecase *BookUsecaseImpl) Create(ctx context.Context, request model.CreateBookRequest) (model.BookResponse, error) {
	// 1. Validasi request
	err := usecase.Validate.Struct(request)
	if err != nil {
		return model.BookResponse{}, err
	}

	// 2. Transaksi DB
	tx, err := usecase.DB.Begin()
	if err != nil {
		return model.BookResponse{}, err
	}
	defer tx.Rollback()

	// 3. Request -> Entity
	book := entity.Book{
		Name:            request.Name,
		Author:          request.Author,
		PublicationYear: request.PublicationYear,
		Description:     request.Description,
	}

	// 4. Simpan ke DB lewat Repository
	result, err := usecase.BookRepository.Create(ctx, tx, &book)
	if err != nil {
		return model.BookResponse{}, err
	}

	// 5. Commit Transaksi
	err = tx.Commit()
	if err != nil {
		return model.BookResponse{}, err
	}

	// 6. Entity -> DTO Response
	return converter.ToBookConverter(result), nil
}

func (usecase *BookUsecaseImpl) Update(ctx context.Context, request model.UpdateBookRequest) (model.BookResponse, error) {
	// 1. Validasi request
	err := usecase.Validate.Struct(request)
	if err != nil {
		return model.BookResponse{}, err
	}

	// 2. Transaksi DB
	tx, err := usecase.DB.Begin()
	if err != nil {
		return model.BookResponse{}, err
	}
	defer tx.Rollback()

	// 3. Cek apakah buku dengan ID tersebut ada di DB
	book, err := usecase.BookRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		return model.BookResponse{}, err
	}

	// 4. Update data entity
	book.Name = request.Name
	book.Author = request.Author
	book.PublicationYear = request.PublicationYear
	book.Description = request.Description

	// 5. Simpan perubahan ke DB
	result, err := usecase.BookRepository.Update(ctx, tx, book)
	if err != nil {
		return model.BookResponse{}, err
	}

	// 6. Commit Transaksi
	err = tx.Commit()
	if err != nil {
		return model.BookResponse{}, err
	}

	return converter.ToBookConverter(result), nil
}

func (usecase *BookUsecaseImpl) Delete(ctx context.Context, bookId int) error {
	tx, err := usecase.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Cek apakah buku ada
	book, err := usecase.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		return err
	}

	// Hapus buku
	err = usecase.BookRepository.Delete(ctx, tx, book)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (usecase *BookUsecaseImpl) FindById(ctx context.Context, bookId int) (model.BookResponse, error) {
	tx, err := usecase.DB.Begin()
	if err != nil {
		return model.BookResponse{}, err
	}
	defer tx.Rollback()

	book, err := usecase.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		return model.BookResponse{}, err
	}

	err = tx.Commit()
	if err != nil {
		return model.BookResponse{}, err
	}

	return converter.ToBookConverter(book), nil
}

func (usecase *BookUsecaseImpl) FindAll(ctx context.Context) ([]model.BookResponse, error) {
	tx, err := usecase.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	books, err := usecase.BookRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// Loop list entity.Book -> convert ke []model.BookResponse
	var bookResponses []model.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, converter.ToBookConverter(&book))
	}

	return bookResponses, nil
}
