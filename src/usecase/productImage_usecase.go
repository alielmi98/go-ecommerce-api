package usecase

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/domain/filter"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type ProductImageUsecase struct {
	base *BaseUsecase[model.ProductImage, dto.CreateProductImage, dto.UpdateProductImage, dto.ResponseProductImage]
}

func NewProductImageUsecase(cfg *config.Config, repository repository.ProductImageRepository) *ProductImageUsecase {
	return &ProductImageUsecase{
		base: NewBaseUsecase[model.ProductImage, dto.CreateProductImage, dto.UpdateProductImage, dto.ResponseProductImage](cfg, repository),
	}
}

// Create a new product image
func (u *ProductImageUsecase) Create(ctx context.Context, req dto.CreateProductImage) (dto.ResponseProductImage, error) {
	return u.base.Create(ctx, req)
}

// Update an existing product image
func (u *ProductImageUsecase) Update(ctx context.Context, id int, req dto.UpdateProductImage) (dto.ResponseProductImage, error) {
	// Pass request to Update method
	return u.base.Update(ctx, id, req)
}

// Delete a product image by ID
func (u *ProductImageUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get a product image by ID
func (u *ProductImageUsecase) GetById(ctx context.Context, id int) (dto.ResponseProductImage, error) {
	return u.base.GetById(ctx, id)
}

// Get product images with filters
func (s *ProductImageUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.ResponseProductImage], error) {
	return s.base.GetByFilter(ctx, req)
}
