package dto

// Payment
type ResponsePayment struct {
	Id          int
	Amount      float64
	Status      string
	AuthorityId string
	RefId       int
	UserId      int
	OrderId     int
}
type ResponsePaymentUrl struct {
	PaymentUrl string
}

type PaymentVerificationResponse struct {
	RefId     int
	PaymentId int
	Status    string
}

type CreatePaymentUrl struct {
	OrderId int
}

// CreatePayment is used to create a new payment
type CreatePayment struct {
	Amount      float64
	Status      string
	AuthorityId string
	RefId       int
	UserId      int
	OrderId     int
}
type UpdatePayment struct {
	Amount      float64
	Status      string
	AuthorityId string
	RefId       int
	UserId      int
	OrderId     int
}

// Order
type ResponseOrder struct {
	Id         int
	UserId     int
	TotalPrice float64
	TotalItems int
	Status     string
	Address    string
	PaymentId  int
	OrderItems []ResponseOrderItem
}

type CreateOrder struct {
	UserId     int
	TotalPrice float64
	TotalItems int
	Status     string
	Address    string
	PaymentId  int
}

type UpdateOrder struct {
	UserId     int
	TotalPrice float64
	TotalItems int
	Status     string
	Address    string
	PaymentId  int
}
type ResponseOrderItem struct {
	ProductId int
	Quantity  int
	OrderId   int
	UnitPrice float64
}

type CreateOrderItem struct {
	ProductId int
	Quantity  int
	OrderId   int
	UnitPrice float64
}

type UpdateOrderItem struct {
	ProductId int
	Quantity  int
	OrderId   int
	UnitPrice float64
}

// Cart
type ResponseCart struct {
	Id         int
	UserId     int
	TotalPrice float64
	TotalItems int
	CartItems  []ResponseCartItem
}
type CreateCart struct {
	UserId int
}

type UpdateCart struct {
	TotalPrice float64
	TotalItems int
	Items      []ResponseCartItem
}

type ResponseCartItem struct {
	ProductId int
	Quantity  int
	CartId    int
	UnitPrice float64
}

type CreateCartItem struct {
	ProductId int
	CartId    int
	Quantity  int
	UnitPrice float64
}

type UpdateCartItem struct {
	Quantity int
}

type CheckOutRequest struct {
	Address string
}
