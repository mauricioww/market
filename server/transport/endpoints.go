package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mauricioww/market/server/service"
)

type Endpoints struct {
	CreateProduct endpoint.Endpoint
}

func MakeEndpoints(s service.Servicer) Endpoints {
	return Endpoints{
		CreateProduct: makeCreateProductEndpoint(s),
	}
}

func makeCreateProductEndpoint(s service.Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProductRequest)
		res, err := s.CreateProduct(ctx, req.Name, req.SupplierId, req.CategoryId, req.UnitsInStock, req.UnitPrice, req.Discontinued)
		return CreateProductResponse{Id: res, Name: req.Name, SupplierId: req.SupplierId, CategoryId: req.CategoryId, UnitsInStock: req.UnitsInStock, UnitPrice: req.UnitPrice, Discontinued: req.Discontinued}, err
	}
}
