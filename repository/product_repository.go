package repository

import (
	"context"
	"database/sql"

	"golang_clean_architecture/entity"
)

type ProductRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Product
	FindById(ctx context.Context, tx *sql.Tx, id int) (entity.Product, error)
	Create(ctx context.Context, tx *sql.Tx, product entity.Product) (entity.Product, error)
}
