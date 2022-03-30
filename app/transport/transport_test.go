package transport_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mauricioww/market/app/errors"
	"github.com/mauricioww/market/app/models"
	"github.com/mauricioww/market/app/transport"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	srvMock := new(transport.ServiceMock)
	endpoints := transport.MakeEndpoints(srvMock)
	s := transport.NewHttpServer(endpoints)
	server := httptest.NewServer(s)

	defer server.Close()

	testCases := []struct {
		testName   string
		body       string
		data       transport.CreateProductRequest
		res        int
		err        error
		httpStatus int
	}{
		{
			testName: "product created success",
			body: ` 
				{
					"name":           "simple product",
					"units_in_stock": 2,
					"unit_price":     100.50
				}
			`,
			data: transport.CreateProductRequest{
				Name:         "simple product",
				UnitsInStock: 2,
				UnitPrice:    100.50,
			},
			res:        1,
			err:        nil,
			httpStatus: 200,
		},
		{
			testName: "no field name error",
			body: `
				{
					"units_in_stock": 2,
					"unit_price":     100.50
				}
			`,
			data: transport.CreateProductRequest{
				UnitsInStock: 2,
				UnitPrice:    100.50,
			},
			res:        -1,
			err:        errors.NewBadRequestNameError(),
			httpStatus: 400,
		},
		{
			testName: "no field units_in_stock error",
			body: `
				{
					"name":           "simple product",
					"unit_price":     100.50
				}
			`,
			data: transport.CreateProductRequest{
				Name:      "simple product",
				UnitPrice: 100.50,
			},
			res:        -1,
			err:        errors.NewBadRequestUnitsInStockError(),
			httpStatus: 400,
		},
		{
			testName: "no field unit_price error",
			body: `
				{
					"name":           "simple product",
					"units_in_stock": 2
				}
			`,
			data: transport.CreateProductRequest{
				Name:         "simple product",
				UnitsInStock: 2,
			},
			res:        -1,
			err:        errors.NewBadRequestUnitPriceError(),
			httpStatus: 400,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert := assert.New(t)

			srvMock.On("CreateProduct", tc.data.Name, tc.data.SupplierId, tc.data.CategoryId,
				tc.data.UnitsInStock, tc.data.UnitPrice, tc.data.Discontinued).Return(tc.res, tc.err)
			res, _ := http.Post(server.URL+"/api/products", "application/json", strings.NewReader(tc.body))

			assert.Equal(tc.httpStatus, res.StatusCode)
		})
	}
}

func TestGetProduct(t *testing.T) {
	srvMock := new(transport.ServiceMock)
	endpoints := transport.MakeEndpoints(srvMock)
	s := transport.NewHttpServer(endpoints)
	server := httptest.NewServer(s)

	defer server.Close()

	testCases := []struct {
		testName   string
		data       uint
		res        models.Product
		err        error
		httpStatus int
	}{
		{
			testName: "product found success",
			data:     1,
			res: models.Product{
				Name:         "fake item",
				SupplierId:   2,
				UnitsInStock: 10,
				UnitPrice:    100.3,
			},
			err:        nil,
			httpStatus: 200,
		},
		{
			testName:   "product not found error",
			data:       0,
			err:        errors.NewProductNotFoundError(),
			httpStatus: 404,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert := assert.New(t)

			srvMock.On("GetProduct", tc.data).Return(tc.res, tc.err)
			uri := fmt.Sprintf("%v/api/products/%v", server.URL, tc.data)
			res, _ := http.Get(uri)

			assert.Equal(res.StatusCode, tc.httpStatus)
		})
	}
}
