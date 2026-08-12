package repository

import (
	"context"
	"database/sql"
	"errors"

	"golang_clean_architecture/internal/entity"
)

type BookRepositoryImpl struct{}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, book *entity.Book) (*entity.Book, error) {
	SQL := "INSERT INTO books(name, author, publication_year, description) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, book.Name, book.Author, book.PublicationYear, book.Description)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	book.ID = int(id)
	return book, nil
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book *entity.Book) (*entity.Book, error) {
	SQL := "UPDATE books SET name = ?, author = ?, publication_year = ?, description = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Name, book.Author, book.PublicationYear, book.Description, book.ID)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, book *entity.Book) error {
	SQL := "DELETE FROM books WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId int) (*entity.Book, error) {
	SQL := "SELECT id, name, author, publication_year, description FROM books WHERE id = ?"

	book := &entity.Book{}
	err := tx.QueryRowContext(ctx, SQL, bookId).Scan(
		&book.ID,
		&book.Name,
		&book.Author,
		&book.PublicationYear,
		&book.Description,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Books tidak ditemukan")
		}
		return nil, err
	}
	return book, nil
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Book, error) {
	SQL := "SELECT id, name, author, publication_year, description FROM books"

	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []entity.Book{}
	for rows.Next() {
		book := entity.Book{}
		err := rows.Scan(
			&book.ID,
			&book.Name,
			&book.Author,
			&book.PublicationYear,
			&book.Description,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
