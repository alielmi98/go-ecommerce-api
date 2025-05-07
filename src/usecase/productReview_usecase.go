package usecase

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/domain/filter"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type ProductReviewUsecase struct {
	base *BaseUsecase[model.ProductReview, dto.CreateProductReview, dto.UpdateProductReview, dto.ResponseProductReview]
}

func NewProductReviewUsecase(cfg *config.Config, repository repository.ProductReviewRepository) *ProductReviewUsecase {
	return &ProductReviewUsecase{
		base: NewBaseUsecase[model.ProductReview, dto.CreateProductReview, dto.UpdateProductReview, dto.ResponseProductReview](cfg, repository),
	}
}

// Create
func (u *ProductReviewUsecase) Create(ctx context.Context, req dto.CreateProductReview) (dto.ResponseProductReview, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *ProductReviewUsecase) Update(ctx context.Context, id int, req dto.UpdateProductReview) (dto.ResponseProductReview, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *ProductReviewUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// GetById
func (u *ProductReviewUsecase) GetById(ctx context.Context, id int) (dto.ResponseProductReview, error) {
	return u.base.GetById(ctx, id)
}

// GetByFilter

func (s *ProductReviewUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.ResponseProductReview], error) {
	return s.base.GetByFilter(ctx, req)
}
