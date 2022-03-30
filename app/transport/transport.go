package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	httpGokit "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/mauricioww/market/app/errors"
)

func NewHttpServer(endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware)

	opt := httpGokit.ServerOption(httpGokit.ServerErrorEncoder(encodeError))

	r.Methods("POST").Path("/api/products").Handler(httpGokit.NewServer(
		endpoints.CreateProduct,
		decodeCreateProductRequest,
		encodeResponse,
		opt,
	))

	r.Methods("GET").Path("/api/products/{id}").Handler(httpGokit.NewServer(
		endpoints.GetProduct,
		decodeGetProductRequest,
		encodeResponse,
		opt,
	))

	r.Methods("PUT").Path("/api/products/{id}").Handler(httpGokit.NewServer(
		endpoints.UpdateProduct,
		decodeUpdateProductRequest,
		encodeResponse,
		opt,
	))

	r.Methods("DELETE").Path("/api/products/{id}").Handler(httpGokit.NewServer(
		endpoints.DeleteProduct,
		decodeDeleteUserRequest,
		encodeResponse,
		opt,
	))

	return r
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func encodeError(_ context.Context, err error, rw http.ResponseWriter) {
	r, ok := err.(errors.Resolver)
	if !ok {
		r = errors.NewUnknownError()
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(errors.ResolveHttp(r.ResolveCode()))
	json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
}

func encodeResponse(ctx context.Context, rw http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(rw).Encode(response)
}

func decodeCreateProductRequest(ctx context.Context, rw *http.Request) (interface{}, error) {
	var req CreateProductRequest
	err := json.NewDecoder(rw.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetProductRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return nil, err
	}

	request := GetProductRequest{Id: uint(id)}
	return request, nil
}

func decodeUpdateProductRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateProductRequest
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.Id = uint(id)
	return req, nil
}

func decodeDeleteUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return nil, err
	}

	request := DeleteProductRequest{Id: uint(id)}
	return request, nil
}
