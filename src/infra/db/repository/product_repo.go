package repository

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/infra/db"
	"github.com/alielmi98/go-ecommerce-api/pkg/service_errors"
)

type PostgresProductRepository struct {
	*BaseRepository[models.Product]
}

func NewProductRepository(preloads []db.PreloadEntity) *PostgresProductRepository {
	return &PostgresProductRepository{BaseRepository: NewBaseRepository[models.Product](preloads)}
}

func (r *PostgresProductRepository) CheckProductAvailability(productId int, orderQuantity int) bool {
	var product models.Product
	err := r.database.Where("id = ?", productId).First(&product).Error
	if err != nil {
		return false
	}
	return product.Stock >= orderQuantity
}

// Deducts the stock of a product by the specified quantity.
func (r *PostgresProductRepository) DeductProductStock(productId int, quantity int) error {
	var product models.Product
	tx := r.database.Begin()
	err := tx.Where("id = ?", productId).First(&product).Error
	if err != nil {
		tx.Rollback()
		return &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
	}
	if product.Stock < quantity {
		tx.Rollback()
		return &service_errors.ServiceError{EndUserMessage: service_errors.InsufficientStock}
	}
	product.Stock -= quantity
	err = tx.Save(&product).Error
	if err != nil {
		tx.Rollback()
		return &service_errors.ServiceError{EndUserMessage: service_errors.UnExpectedError}
	}
	tx.Commit()
	return nil
}

// Increaments the view count of a product by 1.
func (r *PostgresProductRepository) IncrementProductViewCount(productId int) error {
	var product models.Product
	tx := r.database.Begin()
	err := tx.Where("id = ?", productId).First(&product).Error
	if err != nil {
		tx.Rollback()
		return &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
	}
	product.CountViews++
	err = tx.Save(&product).Error
	if err != nil {
		tx.Rollback()
		return &service_errors.ServiceError{EndUserMessage: service_errors.UnExpectedError}
	}
	tx.Commit()
	return nil
}
func (r *PostgresProductRepository) UpdateAverageRating(ctx context.Context, productId int, avg float64) error {
	return r.database.Model(&models.Product{}).
		Where("id = ?", productId).
		Update("average_rating", avg).Error
}
