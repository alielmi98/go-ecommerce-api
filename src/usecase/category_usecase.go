package usecase

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/config"
	model "github.com/alielmi98/go-ecommerce-api/domin/models"
	"github.com/alielmi98/go-ecommerce-api/domin/repository"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type CatecoryUsecase struct {
	base *BaseUsecase[model.Category, dto.CreateCategory, dto.UpdateCategory, dto.ResponseCategory]
}

func NewCategoryUsecase(cfg *config.Config, repository repository.CategoryRepository) *CatecoryUsecase {
	return &CatecoryUsecase{
		base: NewBaseUsecase[model.Category, dto.CreateCategory, dto.UpdateCategory, dto.ResponseCategory](cfg, repository),
	}
}

// Create
func (u *CatecoryUsecase) Create(ctx context.Context, req dto.CreateCategory) (dto.ResponseCategory, error) {
	return u.base.Create(ctx, req)
}
