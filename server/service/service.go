package service

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/mauricioww/market/server/repository"
)

type Service struct {
	repository repository.Repositorier
	logger     log.Logger
}

type Servicer interface {
	CreateProduct(ctx context.Context, name string, supplierId uint32, categoryId uint32, unitsInStock uint32, unitPrice float32, discontinued bool) (int32, error)
}

func NewService(r repository.Repositorier, l log.Logger) *Service {
	return &Service{
		repository: r,
		logger:     l,
	}
}

func (s *Service) CreateProduct(ctx context.Context, name string, supplierId uint32, categoryId uint32, unitInStock uint32, unitPrice float32, discontinued bool) (int32, error) {
	l := log.With(s.logger, "method", "create_product")

	// Do validations
	id, err := s.repository.CreateProduct(ctx, name, supplierId, categoryId, unitInStock, unitPrice, discontinued)

	if err != nil {
		level.Error(l).Log("ERROR", err)
		return -1, err
	}

	l.Log("action", "success")
	return id, nil
}
