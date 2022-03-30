package transport

type CreateProductRequest struct {
	Name         string  `json:"name,omitempty"`
	SupplierId   uint    `json:"supplier_id,omitempty"`
	CategoryId   uint    `json:"category_id,omitempty"`
	UnitsInStock uint    `json:"units_in_stock,omitempty"`
	UnitPrice    float64 `json:"unit_price,omitempty"`
	Discontinued bool    `json:"discontinued,omitempty"`
}

type GetProductRequest struct {
	Id uint `json:"id,omitempty"`
}
