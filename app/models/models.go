package models

type Product struct {
	Name         string
	SupplierId   uint
	CategoryId   uint
	UnitsInStock uint
	UnitPrice    float64
	Discontinued bool
}
