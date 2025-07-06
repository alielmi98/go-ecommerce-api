package usecase

import (
	"context"
	"log"

	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/domain/filter"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type ProductReviewUsecase struct {
	base        *BaseUsecase[model.ProductReview, dto.CreateProductReview, dto.UpdateProductReview, dto.ResponseProductReview]
	productRepo repository.ProductRepository
}

func NewProductReviewUsecase(cfg *config.Config, productReviewRepository repository.ProductReviewRepository, productRepo repository.ProductRepository) *ProductReviewUsecase {
	return &ProductReviewUsecase{
		base:        NewBaseUsecase[model.ProductReview, dto.CreateProductReview, dto.UpdateProductReview, dto.ResponseProductReview](cfg, productReviewRepository),
		productRepo: productRepo,
	}
}

// Create
func (u *ProductReviewUsecase) Create(ctx context.Context, req dto.CreateProductReview) (dto.ResponseProductReview, error) {
	response, err := u.base.Create(ctx, req)
	if err != nil {
		return response, err
	}
	go func() {
		// Update the average rating for the product
		err = u.UpdateProductAverageRating(req.ProductId)
		if err != nil {
			log.Printf("caller:%s  Level:%s Msg:%v", constants.Postgres, constants.UseCase, err)
		}
	}()
	return response, nil
}

// Update
func (u *ProductReviewUsecase) Update(ctx context.Context, id int, req dto.UpdateProductReview) (dto.ResponseProductReview, error) {
	response, err := u.base.Update(ctx, id, req)
	if err != nil {
		return response, err
	}
	go func() {
		// Update the average rating for the product
		err = u.UpdateProductAverageRating(id)
		if err != nil {
			log.Printf("caller:%s  Level:%s Msg:%v", constants.Postgres, constants.UseCase, err)
		}
	}()
	return response, nil
}

// Delete
func (u *ProductReviewUsecase) Delete(ctx context.Context, id int) error {
	err := u.base.Delete(ctx, id)
	if err != nil {
		return err
	}
	go func() {
		// Update the average rating for the product
		err = u.UpdateProductAverageRating(id)
		if err != nil {
			log.Printf("caller:%s  Level:%s Msg:%v", constants.Postgres, constants.UseCase, err)
		}
	}()
	return nil
}

// GetById
func (u *ProductReviewUsecase) GetById(ctx context.Context, id int) (dto.ResponseProductReview, error) {
	return u.base.GetById(ctx, id)
}

// GetByFilter

func (s *ProductReviewUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.ResponseProductReview], error) {
	return s.base.GetByFilter(ctx, req)
}

// UpdateProductAverageRating updates the average rating for a product
func (u *ProductReviewUsecase) UpdateProductAverageRating(productId int) error {
	product, err := u.productRepo.GetById(context.Background(), productId)
	if err != nil {
		return err
	}

	if len(product.Reviews) == 0 {
		return nil
	}

	var totalRating float64
	for _, review := range product.Reviews {
		totalRating += float64(review.Rating)
	}

	averageRating := totalRating / float64(len(product.Reviews))
	return u.productRepo.UpdateAverageRating(context.Background(), productId, averageRating)
}
