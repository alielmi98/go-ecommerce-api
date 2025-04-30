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
