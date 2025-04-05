package usecase

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/domin/filter"
	model "github.com/alielmi98/go-ecommerce-api/domin/models"
	"github.com/alielmi98/go-ecommerce-api/domin/repository"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type CategoryUsecase struct {
	base *BaseUsecase[model.Category, dto.CreateCategory, dto.UpdateCategory, dto.ResponseCategory]
}

func NewCategoryUsecase(cfg *config.Config, repository repository.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{
		base: NewBaseUsecase[model.Category, dto.CreateCategory, dto.UpdateCategory, dto.ResponseCategory](cfg, repository),
	}
}

// Create
func (u *CategoryUsecase) Create(ctx context.Context, req dto.CreateCategory) (dto.ResponseCategory, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *CategoryUsecase) Update(ctx context.Context, id int, req dto.UpdateCategory) (dto.ResponseCategory, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *CategoryUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// GetById
func (u *CategoryUsecase) GetById(ctx context.Context, id int) (dto.ResponseCategory, error) {
	return u.base.GetById(ctx, id)
}

// GetByFilter

func (s *CategoryUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.ResponseCategory], error) {
	return s.base.GetByFilter(ctx, req)
}
