package usecase

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/domain/filter"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type OrderUsecase struct {
	base *BaseUsecase[model.Order, dto.CreateOrder, dto.UpdateOrder, dto.ResponseOrder]
	cfg  *config.Config
}

func NewOrderUsecase(cfg *config.Config, repository repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{
		base: NewBaseUsecase[model.Order, dto.CreateOrder, dto.UpdateOrder, dto.ResponseOrder](cfg, repository),
		cfg:  cfg,
	}
}

// Create a new order
func (u *OrderUsecase) Create(ctx context.Context, req dto.CreateOrder) (dto.ResponseOrder, error) {
	return u.base.Create(ctx, req)
}

// Update an existing order
func (u *OrderUsecase) Update(ctx context.Context, id int, req dto.UpdateOrder) (dto.ResponseOrder, error) {
	return u.base.Update(ctx, id, req)
}

// Delete an order by ID
func (u *OrderUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get an order by ID
func (u *OrderUsecase) GetById(ctx context.Context, id int) (dto.ResponseOrder, error) {
	return u.base.GetById(ctx, id)
}

// Get orders by filter
func (u *OrderUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.ResponseOrder], error) {
	return u.base.GetByFilter(ctx, req)
}
