package repository

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/domin/filter"
	model "github.com/alielmi98/go-ecommerce-api/domin/models"
)

type BaseRepository[TEntity any] interface {
	Create(ctx context.Context, entity TEntity) (TEntity, error)
	Update(ctx context.Context, id int, entity map[string]interface{}) (TEntity, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (TEntity, error)
	GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]TEntity, error)
}
type ProductRepository interface {
	BaseRepository[model.Product]
}

type CategoryRepository interface {
	BaseRepository[model.Category]
}

type ProductImageRepository interface {
	BaseRepository[model.ProductImage]
}

type ProductReviewRepository interface {
	BaseRepository[model.ProductReview]
}
