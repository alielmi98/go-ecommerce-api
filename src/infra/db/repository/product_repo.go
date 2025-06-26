package repository

import (
	"github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/infra/db"
)

type PostgresProductRepository struct {
	*BaseRepository[models.Product]
}

func NewProductRepository(preloads []db.PreloadEntity) *PostgresProductRepository {
	return &PostgresProductRepository{BaseRepository: NewBaseRepository[models.Product](preloads)}
}

func (r *PostgresProductRepository) CheckProductAvailability(productId int, orderQuantity int) bool {
	var product models.Product
	database := db.Preload(r.database, r.preloads)
	err := database.Where("id = ?", productId).First(&product).Error
	if err != nil {
		return false
	}
	return product.Stock >= orderQuantity
}
