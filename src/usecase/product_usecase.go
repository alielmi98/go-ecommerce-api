package usecase

import (
	"context"
	"log"

	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/domain/filter"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/domain/repository"
	"github.com/alielmi98/go-ecommerce-api/events"
	"github.com/alielmi98/go-ecommerce-api/usecase/dto"
)

type ProductUsecase struct {
	base       *BaseUsecase[model.Product, dto.CreateProduct, dto.UpdateProduct, dto.ResponseProduct]
	dispatcher *events.EventDispatcher
	repository repository.ProductRepository
}

func NewProductUsecase(cfg *config.Config, repository repository.ProductRepository, dispatcher *events.EventDispatcher) *ProductUsecase {
	return &ProductUsecase{
		base:       NewBaseUsecase[model.Product, dto.CreateProduct, dto.UpdateProduct, dto.ResponseProduct](cfg, repository),
		dispatcher: dispatcher,
		repository: repository,
	}
}

// Create
func (u *ProductUsecase) Create(ctx context.Context, req dto.CreateProduct) (dto.ResponseProduct, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *ProductUsecase) Update(ctx context.Context, id int, req dto.UpdateProduct) (dto.ResponseProduct, error) {

	if req.Price != 0 {
		currentProduct, err := u.base.repository.GetById(ctx, id)
		if err != nil {
			return dto.ResponseProduct{}, err
		}
		// check if the price has changed
		if currentProduct.Price != req.Price {
			// dispatch the event
			event := events.ProductPriceChangedEvent{
				ProductID: id,
				NewPrice:  req.Price,
			}
			u.dispatcher.Dispatch("ProductPriceChanged", event)
		}
	}

	updatedProduct, err := u.base.Update(ctx, id, req)
	if err != nil {
		return dto.ResponseProduct{}, err
	}
	return updatedProduct, nil
}

// Delete
func (u *ProductUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// GetById
func (u *ProductUsecase) GetById(ctx context.Context, id int) (dto.ResponseProduct, error) {
	// Increment the view count of the product
	go func() {
		err := u.repository.IncrementProductViewCount(id)
		if err != nil {
			log.Printf("caller:%s  Level:%s Msg:%v", constants.Postgres, constants.UseCase, err)
		}
	}()
	return u.base.GetById(ctx, id)
}

// GetByFilter

func (s *ProductUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.ResponseProduct], error) {
	return s.base.GetByFilter(ctx, req)
}
