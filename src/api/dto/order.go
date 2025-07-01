package dto

import (
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

// Cart
type CreateCartItem struct {
	ProductId int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type UpdateCartItem struct {
	Quantity int `json:"quantity" validate:"required"`
}

type CartItemResponse struct {
	ProductId int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

type CreateCart struct {
	UserId int `json:"user_id" validate:"required"`
}

type UpdateCart struct {
	TotalPrice float64 `json:"total_price" validate:"required"`
	TotalItems int     `json:"total_items" validate:"required"`
}
type CartResponse struct {
	Id         int                `json:"id"`
	UserId     int                `json:"user_id"`
	TotalPrice float64            `json:"total_price"`
	TotalItems int                `json:"total_items"`
	CartItems  []CartItemResponse `json:"cart_items"`
}

type CreatePayment struct {
	OrderId int `json:"order_id" validate:"required"`
}

type PaymentResponse struct {
	PaymentUrl string `json:"payment_url"`
}

type UpdatePayment struct {
	Amount      float64 `json:"amount" validate:"required"`
	Status      string  `json:"status" validate:"required"`
	AuthorityId string  `json:"authority_id" validate:"required"`
	RefId       int     `json:"ref_id" validate:"required"`
	UserId      int     `json:"user_id" validate:"required"`
	OrderId     int     `json:"order_id" validate:"required"`
}

func ToPaymentResponse(from dto.ResponsePayment) PaymentResponse {
	return PaymentResponse{
		PaymentUrl: from.PaymentUrl,
	}
}
func ToCreatePayment(from CreatePayment) dto.CreatePayment {
	return dto.CreatePayment{
		OrderId: from.OrderId,
	}
}
func ToUpdatePayment(from UpdatePayment) dto.UpdatePayment {
	return dto.UpdatePayment{
		Amount:      from.Amount,
		Status:      from.Status,
		AuthorityId: from.AuthorityId,
		RefId:       from.RefId,
		UserId:      from.UserId,
		OrderId:     from.OrderId,
	}
}

func ToCartResponse(from dto.ResponseCart) CartResponse {
	return CartResponse{
		Id:         from.Id,
		UserId:     from.UserId,
		TotalPrice: from.TotalPrice,
		TotalItems: from.TotalItems,
		CartItems:  ToCartItemResponseList(from.CartItems),
	}
}
func ToCartItemResponseList(from []dto.ResponseCartItem) []CartItemResponse {
	var cartItems []CartItemResponse
	for _, item := range from {
		cartItems = append(cartItems, CartItemResponse{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		})
	}
	return cartItems
}
func ToCreateCart(from CreateCart) dto.CreateCart {
	return dto.CreateCart{
		UserId: from.UserId,
	}
}
func ToUpdateCart(from UpdateCart) dto.UpdateCart {
	return dto.UpdateCart{
		TotalPrice: from.TotalPrice,
		TotalItems: from.TotalItems,
	}
}

func ToCartItemResponse(from dto.ResponseCartItem) CartItemResponse {
	return CartItemResponse{
		ProductId: from.ProductId,
		Quantity:  from.Quantity,
		UnitPrice: from.UnitPrice,
	}
}

func ToCreateCartItem(from CreateCartItem) dto.CreateCartItem {
	return dto.CreateCartItem{
		ProductId: from.ProductId,
		Quantity:  from.Quantity,
	}
}

func ToUpdateCartItem(from UpdateCartItem) dto.UpdateCartItem {
	return dto.UpdateCartItem{
		Quantity: from.Quantity,
	}
}

// Order

type CreateOrder struct {
	UserId     int     `json:"user_id" validate:"required"`
	TotalPrice float64 `json:"total_price" validate:"required"`
	TotalItems int     `json:"total_items" validate:"required"`
	Status     string  `json:"status" validate:"required"`
	Address    string  `json:"address" validate:"required"`
	PaymentId  int     `json:"payment_id" validate:"required"`
}
type UpdateOrder struct {
	UserId     int     `json:"user_id" validate:"required"`
	TotalPrice float64 `json:"total_price" validate:"required"`
	TotalItems int     `json:"total_items" validate:"required"`
	Status     string  `json:"status" validate:"required"`
	Address    string  `json:"address" validate:"required"`
	PaymentId  int     `json:"payment_id" validate:"required"`
}
type OrderResponse struct {
	Id         int                 `json:"id"`
	UserId     int                 `json:"user_id"`
	TotalPrice float64             `json:"total_price"`
	TotalItems int                 `json:"total_items"`
	Status     string              `json:"status"`
	Address    string              `json:"address"`
	PaymentId  int                 `json:"payment_id"`
	OrderItems []OrderItemResponse `json:"products"`
	CreatedAt  string              `json:"created_at"`
	UpdatedAt  string              `json:"updated_at"`
}
type OrderItemResponse struct {
	ProductId int     `json:"product_id"`
	OrderId   int     `json:"order_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}
type CreateOrderItem struct {
	ProductId int     `json:"product_id" validate:"required"`
	OrderId   int     `json:"order_id" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	UnitPrice float64 `json:"unit_price" validate:"required"`
}
type UpdateOrderItem struct {
	ProductId int     `json:"product_id" validate:"required"`
	OrderId   int     `json:"order_id" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	UnitPrice float64 `json:"unit_price" validate:"required"`
}

type CheckOutRequest struct {
	Address string `json:"address" validate:"required"`
}

func ToOrderResponse(from dto.ResponseOrder) OrderResponse {

	return OrderResponse{
		Id:         from.Id,
		UserId:     from.UserId,
		TotalPrice: from.TotalPrice,
		TotalItems: from.TotalItems,
		Status:     from.Status,
		Address:    from.Address,
		PaymentId:  from.PaymentId,
		OrderItems: ToOrderItemResponseList(from.OrderItems),
	}
}
func ToOrderItemResponseList(from []dto.ResponseOrderItem) []OrderItemResponse {
	var orderItems []OrderItemResponse
	for _, item := range from {
		orderItems = append(orderItems, OrderItemResponse{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		})
	}
	return orderItems
}
func ToCreateOrder(from CreateOrder) dto.CreateOrder {
	return dto.CreateOrder{
		UserId:     from.UserId,
		TotalPrice: from.TotalPrice,
		TotalItems: from.TotalItems,
		Status:     from.Status,
		Address:    from.Address,
		PaymentId:  from.PaymentId,
	}
}
func ToUpdateOrder(from UpdateOrder) dto.UpdateOrder {
	return dto.UpdateOrder{
		UserId:     from.UserId,
		TotalPrice: from.TotalPrice,
		TotalItems: from.TotalItems,
		Status:     from.Status,
		Address:    from.Address,
		PaymentId:  from.PaymentId,
	}
}
func ToOrderItemResponse(from dto.ResponseOrderItem) OrderItemResponse {
	return OrderItemResponse{
		ProductId: from.ProductId,
		OrderId:   from.OrderId,
		Quantity:  from.Quantity,
		UnitPrice: from.UnitPrice,
	}
}
func ToCreateOrderItem(from CreateOrderItem) dto.CreateOrderItem {
	return dto.CreateOrderItem{
		ProductId: from.ProductId,
		OrderId:   from.OrderId,
		Quantity:  from.Quantity,
		UnitPrice: from.UnitPrice,
	}
}
func ToUpdateOrderItem(from UpdateOrderItem) dto.UpdateOrderItem {
	return dto.UpdateOrderItem{
		ProductId: from.ProductId,
		OrderId:   from.OrderId,
		Quantity:  from.Quantity,
		UnitPrice: from.UnitPrice,
	}
}

func ToCheckOutRequest(from CheckOutRequest) dto.CheckOutRequest {
	return dto.CheckOutRequest{
		Address: from.Address,
	}
}
