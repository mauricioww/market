package repository

import (
	"context"
	"database/sql"

	"github.com/go-kit/log"
)

const (
	createProduct = `
		INSERT INTO PRODUCTS(name, supplier_id, category_id, units_in_stock,
			unit_price, discontinued) VALUES (?, ?, ?, ?, ?, ?)
	`

	updateProduct = `
		UPDATE PRODUCTS SET name = ?, supplier_id = ?, category_id = ?, 
			units_in_stock = ?, unit_price = ?, discontinued = ? WHERE id = ?
	`

	getProduct = `
		SELECT name, supplier_id, category_id, units_in_stock, unit_price
			discontinued FROM PRODUCTS WHERE id = ?
	`

	deleteProduct = `
		UPDATE PRODUCTS SET discontinued = true WHERE id = ?
	`
)

type Repository struct {
	db     *sql.DB
	logger log.Logger
}

type Repositorier interface {
	CreateProduct(ctx context.Context, name string, supplierId uint32, categoryId uint32, unitsInStock uint32, unitPrice float32, discontinued bool) (int32, error)
}

func NewRepository(mysql *sql.DB, l log.Logger) *Repository {
	return &Repository{
		db:     mysql,
		logger: log.With(l, "repository", "mysql"),
	}
}

func (r *Repository) CreateProduct(ctx context.Context, name string, supplierId uint32, categoryId uint32, unitsInStock uint32, unitPrice float32, discontinued bool) (int32, error) {
	id, err := r.db.ExecContext(ctx, createProduct, name, supplierId, categoryId, unitsInStock, unitPrice, discontinued)

	if err != nil {
		return -1, err
	}

	n, _ := id.LastInsertId()
	return int32(n), nil
}
