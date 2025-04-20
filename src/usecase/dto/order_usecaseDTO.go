package dto

// Payment
type ResponsePayment struct {
	Id        int
	Amount    float64
	Status    string
	PaymentId string
	UserId    int
	OrderId   int
}

type CreatePayment struct {
	Amount    float64
	Status    string
	PaymentId string
	UserId    int
	OrderId   int
}

type UpdatePayment struct {
	Amount    float64
	Status    string
	PaymentId string
	UserId    int
	OrderId   int
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
	Products   []ResponseOrderItem
}

type CreateOrder struct {
	UserId     int
	TotalPrice float64
	TotalItems int
	Status     string
	Address    string
	PaymentId  int
	Products   []ResponseOrderItem
}

type UpdateOrder struct {
	UserId     int
	TotalPrice float64
	TotalItems int
	Status     string
	Address    string
	PaymentId  int
	Products   []ResponseOrderItem
}
type ResponseOrderItem struct {
	ProductId int
	Quantity  int
	UnitPrice float64
}

type CreateOrderItem struct {
	ProductId int
	Quantity  int
	UnitPrice float64
}

type UpdateOrderItem struct {
	ProductId int
	Quantity  int
	UnitPrice float64
}

// Cart
type ResponseCart struct {
	Id         int
	UserId     int
	TotalPrice float64
	TotalItems int
	Products   []ResponseCartItem
}

type CreateCart struct {
	UserId     int
	TotalPrice float64
	TotalItems int
	Products   []ResponseCartItem
}

type UpdateCart struct {
	UserId     int
	TotalPrice float64
	TotalItems int
	Products   []ResponseCartItem
}

type ResponseCartItem struct {
	ProductId int
	Quantity  int
	UnitPrice float64
}

type CreateCartItem struct {
	ProductId int
	Quantity  int
	UnitPrice float64
}

type UpdateCartItem struct {
	ProductId int
	Quantity  int
	UnitPrice float64
}
