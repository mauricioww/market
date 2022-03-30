package service_test

import (
	"os"
	"testing"

	"github.com/go-kit/log"
	"github.com/mauricioww/market/app/errors"
	"github.com/mauricioww/market/app/models"
	"github.com/mauricioww/market/app/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	repository := new(service.RepositoryMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	srv := service.NewService(repository, logger)

	testCases := []struct {
		testName string
		data     models.Product
		res      int
		err      error
	}{
		{
			testName: "product created success",
			data: models.Product{
				Name:         "simple product",
				UnitsInStock: 10,
				UnitPrice:    100.40,
			},
			res: 1,
			err: nil,
		},
		{
			testName: "no field name error",
			data: models.Product{
				UnitsInStock: 10,
				UnitPrice:    100.40,
			},
			res: -1,
			err: errors.NewBadRequestNameError(),
		},
		{
			testName: "no field units_in_stock error",
			data: models.Product{
				Name:      "simple product",
				UnitPrice: 100.40,
			},
			res: -1,
			err: errors.NewBadRequestUnitsInStockError(),
		},
		{
			testName: "no field unit_price error",
			data: models.Product{
				Name:         "simple product",
				UnitsInStock: 10,
			},
			res: -1,
			err: errors.NewBadRequestUnitPriceError(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert := assert.New(t)

			repository.On("CreateProduct", tc.data.Name, tc.data.SupplierId, tc.data.CategoryId, tc.data.UnitsInStock,
				tc.data.UnitPrice, tc.data.Discontinued).Return(tc.res, tc.err)
			res, err := srv.CreateProduct(tc.data.Name, tc.data.SupplierId, tc.data.CategoryId, tc.data.UnitsInStock, tc.data.UnitPrice, tc.data.Discontinued)

			assert.Equal(int32(tc.res), res)
			assert.Equal(tc.err, err)
		})
	}
}

func TestGetProduct(t *testing.T) {
	repository := new(service.RepositoryMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	srv := service.NewService(repository, logger)

	testCases := []struct {
		testName string
		data     uint
		res      models.Product
		err      error
	}{
		{
			testName: "product found success",
			data:     1,
			res: models.Product{
				Name:         "fake item",
				SupplierId:   12,
				CategoryId:   1,
				UnitsInStock: 100,
				UnitPrice:    100.2,
				Discontinued: false,
			},
			err: nil,
		},
		{
			testName: "product not found error",
			data:     0,
			err:      errors.NewProductNotFoundError(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert := assert.New(t)

			repository.On("GetProduct", tc.data).Return(tc.res, tc.err)
			res, err := srv.GetProduct(tc.data)

			assert.Equal(res, tc.res)
			assert.Equal(err, tc.err)

		})
	}
}
