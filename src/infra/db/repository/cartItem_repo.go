package repository

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/infra/db"
)

type PostgresCartItemRepository struct {
	*BaseRepository[models.CartItem]
}

func NewCartItemRepository(preloads []db.PreloadEntity) *PostgresCartItemRepository {
	return &PostgresCartItemRepository{BaseRepository: NewBaseRepository[models.CartItem](preloads)}
}

// UpdateCartItemsPrice updates the UnitPrice of CartItems for a specific product.
func (r *PostgresCartItemRepository) UpdateCartItemsPrice(ctx context.Context, productId int, newPrice float64) error {
	var cartItems []models.CartItem
	result := r.database.Model(&cartItems).
		Where("product_id = ?", productId).
		Update("unit_price", newPrice)
	return result.Error
}
