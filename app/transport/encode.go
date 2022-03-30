package transport

type CreateProductResponse struct {
	Id           int32   `json:"id"`
	Name         string  `json:"name"`
	SupplierId   uint    `json:"supplier_id"`
	CategoryId   uint    `json:"category_id"`
	UnitsInStock uint    `json:"units_in_stock"`
	UnitPrice    float64 `json:"unit_price"`
	Discontinued bool    `json:"discontinued"`
}

type GetProductResponse struct {
	Name         string  `json:"name"`
	SupplierId   uint    `json:"supplier_id"`
	CategoryId   uint    `json:"category_id"`
	UnitsInStock uint    `json:"units_in_stock"`
	UnitPrice    float64 `json:"unit_price"`
	Discontinued bool    `json:"discontinued"`
}
