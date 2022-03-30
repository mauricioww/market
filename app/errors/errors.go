package errors

const (
	unknownError           = -1
	badRequestName         = 0
	badRequestUnitsInStock = 1
	badRequestUnitPrice    = 2
	productNotFound        = 3
	serverFail             = 4
)

var message = map[int]string{
	unknownError:           "Unsupported error",
	badRequestName:         "Missing field 'name'",
	badRequestUnitsInStock: "Missing field 'units_in_stock'",
	badRequestUnitPrice:    "Missing field 'unit_price'",
	productNotFound:        "Product not found",
	serverFail:             "Internal server error",
}

func messageError(code int) string {
	return message[code]
}

func ResolveHttp(c int) int {
	switch c {
	case badRequestName, badRequestUnitPrice, badRequestUnitsInStock:
		return 400
	case productNotFound:
		return 404
	default:
		return 500
	}
}

type UnknownError int

type BadRequestNameError int

type BadRequestUnitPriceError int

type BadRequestUnitsInStockError int

type ProductNotFoundError int

type ServerFail int

func NewUnknownError() UnknownError {
	return unknownError
}

func NewBadRequestNameError() BadRequestNameError {
	return badRequestName
}

func NewBadRequestUnitPriceError() BadRequestUnitPriceError {
	return badRequestUnitPrice
}

func NewBadRequestUnitsInStockError() BadRequestUnitsInStockError {
	return badRequestUnitsInStock
}

func NewProductNotFoundError() ProductNotFoundError {
	return productNotFound
}

func NewServerFailError() ServerFail {
	return serverFail
}

func (e UnknownError) Error() string {
	return messageError(int(e))
}

func (e BadRequestNameError) Error() string {
	return messageError(int(e))
}

func (e BadRequestUnitPriceError) Error() string {
	return messageError(int(e))
}

func (e BadRequestUnitsInStockError) Error() string {
	return messageError(int(e))
}

func (e ProductNotFoundError) Error() string {
	return messageError(int(e))
}

func (e ServerFail) Error() string {
	return messageError(int(e))
}

type Resolver interface {
	ResolveCode() int
}

func (e UnknownError) ResolveCode() int {
	return int(e)
}

func (e BadRequestNameError) ResolveCode() int {
	return int(e)
}

func (e BadRequestUnitPriceError) ResolveCode() int {
	return int(e)
}

func (e BadRequestUnitsInStockError) ResolveCode() int {
	return int(e)
}

func (e ProductNotFoundError) ResolveCode() int {
	return int(e)
}

func (e ServerFail) ResolveCode() int {
	return int(e)
}
