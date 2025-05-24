package usecase

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/domain/filter"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type OrderItemUsecase struct {
	base *BaseUsecase[model.OrderItem, dto.CreateOrderItem, dto.UpdateOrderItem, dto.ResponseOrderItem]
	cfg  *config.Config
}

func NewOrderItemUsecase(cfg *config.Config, repository repository.OrderItemRepository) *OrderItemUsecase {
	return &OrderItemUsecase{
		base: NewBaseUsecase[model.OrderItem, dto.CreateOrderItem, dto.UpdateOrderItem, dto.ResponseOrderItem](cfg, repository),
		cfg:  cfg,
	}
}

// Create a new order
func (u *OrderItemUsecase) Create(ctx context.Context, req dto.CreateOrderItem) (dto.ResponseOrderItem, error) {
	return u.base.Create(ctx, req)
}

// Update an existing order
func (u *OrderItemUsecase) Update(ctx context.Context, id int, req dto.UpdateOrderItem) (dto.ResponseOrderItem, error) {
	return u.base.Update(ctx, id, req)
}

// Delete an order by ID
func (u *OrderItemUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get an order by ID
func (u *OrderItemUsecase) GetById(ctx context.Context, id int) (dto.ResponseOrderItem, error) {
	return u.base.GetById(ctx, id)
}

// Get orders by filter
func (u *OrderItemUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.ResponseOrderItem], error) {
	return u.base.GetByFilter(ctx, req)
}
