package usecase

import (
	"context"
	"fmt"

	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type CartItemUsecase struct {
	base         *BaseUsecase[model.CartItem, dto.CreateCartItem, dto.UpdateCartItem, dto.ResponseCartItem]
	cartItemRepo repository.CartItemRepository
	cartRepo     repository.CartRepository
	productRepo  repository.ProductRepository
	cfg          *config.Config
}

func NewCartItemUsecase(cfg *config.Config, cartItemRepository repository.CartItemRepository, cartRepository repository.CartRepository, productRepository repository.ProductRepository) *CartItemUsecase {
	return &CartItemUsecase{
		base:         NewBaseUsecase[model.CartItem, dto.CreateCartItem, dto.UpdateCartItem, dto.ResponseCartItem](cfg, cartItemRepository),
		cartItemRepo: cartItemRepository,
		cartRepo:     cartRepository,
		productRepo:  productRepository,
		cfg:          cfg,
	}
}

// Create a new cart item
func (u *CartItemUsecase) Create(ctx context.Context, req dto.CreateCartItem) (dto.ResponseCartItem, error) {
	// Extract the User ID from the request
	userId := int(ctx.Value(constants.UserIdKey).(float64))
	cart, err := u.cartRepo.GetCartByUserId(ctx, userId)
	if err != nil {
		return dto.ResponseCartItem{}, err
	}
	// Check if the cart ID is valid
	if cart == nil {
		return dto.ResponseCartItem{}, fmt.Errorf("cart not found for user ID: %d", userId)
	}

	product, err := u.productRepo.GetById(ctx, req.ProductId)
	if err != nil {
		return dto.ResponseCartItem{}, err
	}

	// Set the Product Price in the request
	req.UnitPrice = product.Price
	// Set the User ID in the request
	req.CartId = cart.Id

	// Check if the product already exists in the cart
	for item := range cart.CartItems {
		if cart.CartItems[item].ProductId == req.ProductId {
			cart.CartItems[item].Quantity += req.Quantity
			// Update the existing cart item with the new quantity
			updatedItem, err := u.base.Update(ctx, cart.CartItems[item].Id, dto.UpdateCartItem{Quantity: cart.CartItems[item].Quantity})
			if err != nil {
				return dto.ResponseCartItem{}, err
			}
			return updatedItem, nil
		}
	}

	// Call the base Create method to create the cart item
	return u.base.Create(ctx, req)

}

// Update a cart item
func (u *CartItemUsecase) Update(ctx context.Context, id int, req dto.UpdateCartItem) (dto.ResponseCartItem, error) {
	return u.base.Update(ctx, id, req)
}

// Delete a cart item
func (u *CartItemUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// GetCartItemsByUserId retrieves all cart items for a specific user
func (u *CartItemUsecase) GetById(ctx context.Context, id int) (dto.ResponseCartItem, error) {
	return u.base.GetById(ctx, id)
}

// UpdateCartItemsPrice updates the UnitPrice of CartItems for a specific product.
func (u *CartItemUsecase) UpdateCartItemsPrice(ctx context.Context, productId int, newPrice float64) error {
	return u.cartItemRepo.UpdateCartItemsPrice(ctx, productId, newPrice)
}
