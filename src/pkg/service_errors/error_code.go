package service_errors

const (
	// Token
	UnExpectedError     = "Expected error"
	ClaimsNotFound      = "Claims not found"
	TokenRequired       = "token required"
	TokenExpired        = "token expired"
	TokenInvalid        = "token invalid"
	InvalidRefreshToken = "invalid refresh token"
	InvalidRolesFormat  = "invalid roles format"
	// User
	EmailExists               = "Email exists"
	UsernameExists            = "Username exists"
	PermissionDenied          = "Permission denied"
	UsernameOrPasswordInvalid = "username or password invalid"

	// DB
	RecordNotFound = "record not found"

	// Order Checkout
	CartIsEmpty         = "the cart is empty"
	ErrItemsUnavailable = "Some items in your cart are unavailable."
)
