package dto

import (
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

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
