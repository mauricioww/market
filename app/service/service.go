package service

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/mauricioww/market/app/errors"
	"github.com/mauricioww/market/app/models"
	"github.com/mauricioww/market/app/repository"
)

type Service struct {
	repository repository.Repositorier
	logger     log.Logger
}

type Servicer interface {
	CreateProduct(name string, supplierId uint, categoryId uint, unitsInStock uint, unitPrice float64, discontinued bool) (int32, error)
	GetProduct(id uint) (models.Product, error)
	UpdateProduct(id uint, name string, supplierId uint, categoryId uint, unitsInStock uint, unitPrice float64, discontinued bool) (bool, error)
	DeleteProduct(id uint) (bool, error)
}

func NewService(r repository.Repositorier, l log.Logger) *Service {
	return &Service{
		repository: r,
		logger:     l,
	}
}

func (s *Service) CreateProduct(name string, supplierId uint, categoryId uint, unitsInStock uint, unitPrice float64, discontinued bool) (int32, error) {
	l := log.With(s.logger, "method", "CreateProduct")

	if name == "" {
		e := errors.NewBadRequestNameError()
		level.Error(l).Log("validation_fail: ", e)
		return -1, e
	}

	if unitsInStock == 0 {
		e := errors.NewBadRequestUnitsInStockError()
		level.Error(l).Log("validation_fail: ", e)
		return -1, e
	}

	if unitPrice == 0 {
		e := errors.NewBadRequestUnitPriceError()
		level.Error(l).Log("validation_fail: ", e)
		return -1, e
	}

	res, err := s.repository.CreateProduct(name, supplierId, categoryId, unitsInStock, unitPrice, discontinued)

	if err != nil {
		level.Error(l).Log("ERROR", err)
		return -1, err
	}

	l.Log("action", "success")
	return res, nil
}

func (s *Service) GetProduct(id uint) (models.Product, error) {
	l := log.With(s.logger, "method", "GetProduct")
	l.Log("id", id)

	res, err := s.repository.GetProduct(id)

	if err != nil {
		level.Error(l).Log("ERROR", err)
		return models.Product{}, err
	}

	l.Log("action", "success")
	return res, nil
}

func (s *Service) UpdateProduct(id uint, name string, supplierId uint, categoryId uint, unitsInStock uint, unitPrice float64, discontinued bool) (bool, error) {
	l := log.With(s.logger, "method", "UpdateProduct")
	l.Log("id", id)

	if name == "" {
		e := errors.NewBadRequestNameError()
		level.Error(l).Log("validation_fail: ", e)
		return false, e
	}

	if unitsInStock == 0 {
		e := errors.NewBadRequestUnitsInStockError()
		level.Error(l).Log("validation_fail: ", e)
		return false, e
	}

	if unitPrice == 0 {
		e := errors.NewBadRequestUnitPriceError()
		level.Error(l).Log("validation_fail: ", e)
		return false, e
	}

	res, err := s.repository.UpdateProduct(id, name, supplierId, categoryId, unitsInStock, unitPrice, discontinued)

	if err != nil {
		level.Error(l).Log("ERROR", err)
		return false, err
	}

	l.Log("action", "success")
	return res, nil
}

func (s *Service) DeleteProduct(id uint) (bool, error) {
	l := log.With(s.logger, "method", "DeleteProduct")
	l.Log("id", id)

	res, err := s.repository.DeleteProduct(id)

	if err != nil {
		level.Error(l).Log("ERROR", err)
		return false, err
	}

	l.Log("action", "success")
	return res, nil
}
