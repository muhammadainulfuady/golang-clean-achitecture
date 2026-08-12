package repository

import (
	"context"
	"database/sql"

	"golang_clean_architecture/internal/entity"
)

type BookRepository interface {
	Create(ctx context.Context, tx *sql.Tx, book *entity.Book) (*entity.Book, error)
	Update(ctx context.Context, tx *sql.Tx, book *entity.Book) (*entity.Book, error)
	Delete(ctx context.Context, tx *sql.Tx, book *entity.Book) error
	FindById(ctx context.Context, tx *sql.Tx, bookId int) (*entity.Book, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Book, error)
}
