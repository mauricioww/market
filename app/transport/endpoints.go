package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mauricioww/market/app/service"
)

type Endpoints struct {
	CreateProduct endpoint.Endpoint
	GetProduct    endpoint.Endpoint
}

func MakeEndpoints(s service.Servicer) Endpoints {
	return Endpoints{
		CreateProduct: makeCreateProductEndpoint(s),
		GetProduct:    makeGetProductEndpoint(s),
	}
}

func makeCreateProductEndpoint(s service.Servicer) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProductRequest)
		res, err := s.CreateProduct(req.Name, req.SupplierId, req.CategoryId, req.UnitsInStock, req.UnitPrice, req.Discontinued)
		return CreateProductResponse{Id: res, Name: req.Name, SupplierId: req.SupplierId, CategoryId: req.CategoryId, UnitsInStock: req.UnitsInStock, UnitPrice: req.UnitPrice, Discontinued: req.Discontinued}, err
	}
}

func makeGetProductEndpoint(s service.Servicer) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductRequest)
		res, err := s.GetProduct(req.Id)
		return GetProductResponse{Name: res.Name, SupplierId: res.SupplierId, CategoryId: res.CategoryId, UnitsInStock: res.UnitsInStock, UnitPrice: res.UnitPrice, Discontinued: res.Discontinued}, err
	}
}
