package usecase

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/common"
	"github.com/alielmi98/go-ecommerce-api/config"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type CartUsecase struct {
	base *BaseUsecase[model.Cart, dto.CreateCart, dto.UpdateCart, dto.ResponseCart]
	repo repository.CartRepository
	cfg  *config.Config
}

func NewCartUsecase(cfg *config.Config, repository repository.CartRepository) *CartUsecase {
	return &CartUsecase{
		base: NewBaseUsecase[model.Cart, dto.CreateCart, dto.UpdateCart, dto.ResponseCart](cfg, repository),
		repo: repository,
		cfg:  cfg,
	}
}

// Create a new cart
func (u *CartUsecase) Create(ctx context.Context, req dto.CreateCart) (dto.ResponseCart, error) {
	return u.base.Create(ctx, req)
}

// Update an existing cart
func (u *CartUsecase) Update(ctx context.Context, id int, req dto.UpdateCart) (dto.ResponseCart, error) {
	// Pass request to Update method
	return u.base.Update(ctx, id, req)
}

// Delete a cart by ID
func (u *CartUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get a cart by ID
func (u *CartUsecase) GetById(ctx context.Context, id int) (dto.ResponseCart, error) {
	return u.base.GetById(ctx, id)
}

// Get cart by user ID
func (u *CartUsecase) GetByUserId(ctx context.Context, userId int) (dto.ResponseCart, error) {
	cart, err := u.repo.GetCartByUserId(ctx, userId)
	if err != nil {
		return dto.ResponseCart{}, err
	}
	if cart == nil {
		return dto.ResponseCart{}, nil
	}
	return common.TypeConverter[dto.ResponseCart](cart)
}
