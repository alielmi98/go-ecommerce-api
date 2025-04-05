package usecase

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/domin/filter"
	model "github.com/alielmi98/go-ecommerce-api/domin/models"
	"github.com/alielmi98/go-ecommerce-api/domin/repository"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type ProductUsecase struct {
	base *BaseUsecase[model.Product, dto.CreateProduct, dto.UpdateProduct, dto.ResponseProduct]
}

func NewProductUsecase(cfg *config.Config, repository repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		base: NewBaseUsecase[model.Product, dto.CreateProduct, dto.UpdateProduct, dto.ResponseProduct](cfg, repository),
	}
}

// Create
func (u *ProductUsecase) Create(ctx context.Context, req dto.CreateProduct) (dto.ResponseProduct, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *ProductUsecase) Update(ctx context.Context, id int, req dto.UpdateProduct) (dto.ResponseProduct, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *ProductUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// GetById
func (u *ProductUsecase) GetById(ctx context.Context, id int) (dto.ResponseProduct, error) {
	return u.base.GetById(ctx, id)
}

// GetByFilter

func (s *ProductUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.ResponseProduct], error) {
	return s.base.GetByFilter(ctx, req)
}
