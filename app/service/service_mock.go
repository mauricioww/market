package service

import (
	"github.com/mauricioww/market/app/models"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) CreateProduct(name string, supplierId uint, categoryId uint, unitsInStock uint, unitPrice float64, discontinued bool) (int32, error) {
	args := r.Called(name, supplierId, categoryId, unitsInStock, unitPrice, discontinued)

	return int32(args.Int(0)), args.Error(1)
}

func (r *RepositoryMock) GetProduct(id uint) (models.Product, error) {
	args := r.Called(id)

	return args.Get(0).(models.Product), args.Error(1)
}

func (r *RepositoryMock) UpdateProduct(id uint, name string, supplierId uint, categoryId uint, unitsInStock uint, unitPrice float64, discontinued bool) (bool, error) {
	args := r.Called(id, name, supplierId, categoryId, unitsInStock, unitPrice, discontinued)

	return args.Bool(0), args.Error(1)
}

func (r *RepositoryMock) DeleteProduct(id uint) (bool, error) {
	args := r.Called(id)

	return args.Bool(0), args.Error(1)
}
