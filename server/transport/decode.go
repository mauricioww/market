package transport

type CreateProductRequest struct {
	Name         string  `json:"name,omitempty"`
	SupplierId   uint32  `json:"supplier_id,omitempty"`
	CategoryId   uint32  `json:"category_id,omitempty"`
	UnitsInStock uint32  `json:"units_in_stock,omitempty"`
	UnitPrice    float32 `json:"unit_price,omitempty"`
	Discontinued bool    `json:"discontinued,omitempty"`
}
