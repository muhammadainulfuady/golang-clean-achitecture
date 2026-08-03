package repository

import (
	"context"
	"database/sql"

	"golang_clean_architecture/entity"
	"golang_clean_architecture/helper"
)

type ProductRepositoryImpl struct{}

func (product *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Product {
	SQL := "SELECT id_product, name, stok, price FROM product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var p entity.Product
		err := rows.Scan(&p.IdProduct, &p.Name, &p.Stok, &p.Price)
		helper.PanicIfError(err)

		products = append(products, p)
	}

	helper.PanicIfError(rows.Err())

	return products
}

func (product *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (entity.Product, error) {
	SQL := "SELECT id_product, name, stok, price FROM product WHERE id_product = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	var p entity.Product
	if rows.Next() {
		err := rows.Scan(&p.IdProduct, &p.Name, &p.Stok, &p.Price)
		helper.PanicIfError(err)
		return p, nil
	}

	return p, sql.ErrNoRows
}

func (product *ProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, p entity.Product) (entity.Product, error) {
	SQL := "INSERT INTO product(name, stok, price) VALUES(?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, p.Name, p.Stok, p.Price)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	p.IdProduct = int(id)
	return p, nil
}
