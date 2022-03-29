package transport

import (
	"context"
	"encoding/json"
	"net/http"

	httpGokit "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware)

	r.Methods("POST").Path("/api/products").Handler(httpGokit.NewServer(
		endpoints.CreateProduct,
		decodeCreateProductRequest,
		encodeResponse,
	))

	return r
}

func decodeCreateProductRequest(ctx context.Context, rw *http.Request) (interface{}, error) {
	var req CreateProductRequest
	err := json.NewDecoder(rw.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, rw http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(rw).Encode(response)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
