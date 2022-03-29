package transport

type CreateProductResponse struct {
	Id           int32   `json:"id"`
	Name         string  `json:"name"`
	SupplierId   uint32  `json:"supplier_id"`
	CategoryId   uint32  `json:"category_id"`
	UnitsInStock uint32  `json:"units_in_stock"`
	UnitPrice    float32 `json:"unit_price"`
	Discontinued bool    `json:"discontinued"`
}
