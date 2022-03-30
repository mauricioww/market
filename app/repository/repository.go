package repository

import (
	"database/sql"

	"github.com/go-kit/log"
	"github.com/mauricioww/market/app/errors"
	"github.com/mauricioww/market/app/models"
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
		SELECT name, supplier_id, category_id, units_in_stock, unit_price,
			discontinued FROM PRODUCTS WHERE id = ?
	`

	deleteProduct = `
		DELETE FROM PRODUCTS WHERE id = ?
	`
)

type Repository struct {
	db     *sql.DB
	logger log.Logger
}

type Repositorier interface {
	CreateProduct(name string, supplierId uint, categoryId uint, unitsInStock uint, unitPrice float64, discontinued bool) (int32, error)
	GetProduct(id uint) (models.Product, error)
	UpdateProduct(id uint, name string, supplierId uint, categoryId uint, unitsInStock uint, unitPrice float64, discontinued bool) (bool, error)
	DeleteProduct(id uint) (bool, error)
}

func NewRepository(mysql *sql.DB, l log.Logger) *Repository {
	return &Repository{
		db:     mysql,
		logger: log.With(l, "repository", "mysql"),
	}
}

func (r *Repository) CreateProduct(name string, supplierId uint, categoryId uint, unitsInStock uint, unitPrice float64, discontinued bool) (int32, error) {
	id, err := r.db.Exec(createProduct, name, supplierId, categoryId, unitsInStock, unitPrice, discontinued)

	if err != nil {
		return -1, errors.NewServerFailError()
	}

	n, _ := id.LastInsertId()
	return int32(n), nil
}

func (r *Repository) GetProduct(id uint) (models.Product, error) {
	var p models.Product
	err := r.db.QueryRow(getProduct, id).Scan(&p.Name, &p.SupplierId, &p.CategoryId, &p.UnitsInStock, &p.UnitPrice, &p.Discontinued)

	if err == sql.ErrNoRows {
		return models.Product{}, errors.NewProductNotFoundError()
	}

	if err != nil {
		return models.Product{}, errors.NewServerFailError()
	}

	return p, nil
}

func (r *Repository) UpdateProduct(id uint, name string, supplierId uint, categoryId uint, unitsInStock uint, unitPrice float64, discontinued bool) (bool, error) {

	if err := r.db.QueryRow(getProduct, id).Scan(); err == sql.ErrNoRows {
		return false, errors.NewProductNotFoundError()
	}

	if _, err := r.db.Exec(updateProduct, name, supplierId, categoryId, unitsInStock, unitPrice, discontinued, id); err != nil {
		return false, errors.NewServerFailError()
	}

	return true, nil
}

func (r *Repository) DeleteProduct(id uint) (bool, error) {

	if err := r.db.QueryRow(getProduct, id).Scan(); err == sql.ErrNoRows {
		return false, errors.NewProductNotFoundError()
	}

	if _, err := r.db.Exec(deleteProduct, id); err != nil {
		return false, errors.NewServerFailError()
	}

	return true, nil
}
