package usecase

import (
	"context"
	"fmt"

	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/domin/models"
	"github.com/alielmi98/go-ecommerce-api/domin/repository"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type CartItemUsecase struct {
	base *BaseUsecase[models.CartItem, dto.CreateCartItem, dto.UpdateCartItem, dto.ResponseCartItem]
	repo repository.CartItemRepository
	cfg  *config.Config
}

func NewCartItemUsecase(cfg *config.Config, repository repository.CartItemRepository) *CartItemUsecase {
	return &CartItemUsecase{
		base: NewBaseUsecase[models.CartItem, dto.CreateCartItem, dto.UpdateCartItem, dto.ResponseCartItem](cfg, repository),
		repo: repository,
		cfg:  cfg,
	}
}

// Create a new cart item
func (u *CartItemUsecase) Create(ctx context.Context, req dto.CreateCartItem) (dto.ResponseCartItem, error) {
	// Extract the User ID from the request
	userId := int(ctx.Value(constants.UserIdKey).(float64))
	cart, err := u.repo.GetCartByUserId(ctx, userId)
	if err != nil {
		return dto.ResponseCartItem{}, err
	}
	// Check if the cart ID is valid
	if cart == nil {
		return dto.ResponseCartItem{}, fmt.Errorf("cart not found for user ID: %d", userId)
	}

	product, err := u.repo.GetCartItemProduct(ctx, req.ProductId)
	if err != nil {
		return dto.ResponseCartItem{}, err
	}

	// Set the Product Price in the request
	req.UnitPrice = product.Price
	// Set the User ID in the request
	req.CartId = cart.Id
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
