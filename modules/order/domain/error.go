package orderdomain

const (
	ErrOrderNotFound       = "order not found"
	ErrUserIdRequired      = "user id is required"
	ErrFoodIdRequired      = "food id is required"
	ErrQuantityRequired    = "quantity is required"
	ErrQuantityInvalid     = "quantity must be greater than 0"
	ErrTotalPriceInvalid   = "total price must be greater than 0"
	ErrCannotCreateOrder   = "cannot create order"
	ErrCannotUpdateOrder   = "cannot update order"
	ErrCannotDeleteOrder   = "cannot delete order"
)

