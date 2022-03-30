package transport

import (
	"github.com/mauricioww/market/app/models"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) CreateProduct(name string, supplierId uint, categoryId uint, unitInStock uint, unitPrice float64, discontinued bool) (int32, error) {
	args := s.Called(name, supplierId, categoryId, unitInStock, unitPrice, discontinued)

	return int32(args.Int(0)), args.Error(1)
}

func (r *ServiceMock) GetProduct(id uint) (models.Product, error) {
	args := r.Called(id)

	return args.Get(0).(models.Product), args.Error(1)
}
