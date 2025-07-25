package repository

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/domain/filter"
	model "github.com/alielmi98/go-ecommerce-api/domain/models"
	"gorm.io/gorm"
)

type BaseRepository[TEntity any] interface {
	Create(ctx context.Context, entity TEntity) (TEntity, error)
	Update(ctx context.Context, id int, entity TEntity) (TEntity, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (TEntity, error)
	GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]TEntity, error)
	BeginTransaction(ctx context.Context) (*gorm.DB, error)
	CreateTx(tx *gorm.DB, entity TEntity) (TEntity, error)
}
type ProductRepository interface {
	BaseRepository[model.Product]
	CheckProductAvailability(ctx context.Context, productId int, orderQuantity int) bool
	DeductProductStock(ctx context.Context, productId int, quantity int) error
	IncrementProductViewCount(ctx context.Context, productId int) error
	UpdateAverageRating(ctx context.Context, productId int, avg float64) error
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

type FileRepository interface {
	BaseRepository[model.File]
}

// Order
type CartRepository interface {
	BaseRepository[model.Cart]
	GetCartByUserId(ctx context.Context, userId int) (*model.Cart, error)
}

type CartItemRepository interface {
	BaseRepository[model.CartItem]
	UpdateCartItemsPrice(ctx context.Context, productId int, newPrice float64) error
}

type OrderRepository interface {
	BaseRepository[model.Order]
}

type OrderItemRepository interface {
	BaseRepository[model.OrderItem]
}

type PaymentRepository interface {
	BaseRepository[model.Payment]
	GetByAuthority(ctx context.Context, authority string) (*model.Payment, error)
}

// Account
type UserRepository interface {
	Create(user *model.User) error
	FindByUsername(username string) (*model.User, error)
}

type RoleRepository interface {
	BaseRepository[model.Role]
}
