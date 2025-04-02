package repository

import (
	"context"

	model "github.com/alielmi98/go-ecommerce-api/domin/models"
)

type BaseRepository[TEntity any] interface {
	Create(ctx context.Context, entity TEntity) (TEntity, error)
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
