package models

type Product struct {
	Id           uint
	Name         string
	SupplierId   uint
	CategoryId   uint
	UnitsInStock uint
	UnitPrice    float64
	Discontinued bool
}
