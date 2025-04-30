package repository

import (
	"context"
	"log"

	"github.com/alielmi98/go-ecommerce-api/constants"

	"github.com/alielmi98/go-ecommerce-api/domin/models"
	"github.com/alielmi98/go-ecommerce-api/infra/db"
	"github.com/alielmi98/go-ecommerce-api/pkg/service_errors"
	"gorm.io/gorm"
)

type PostgresCartItemRepository struct {
	*BaseRepository[models.CartItem]
}

func NewCartItemRepository(preloads []db.PreloadEntity) *PostgresCartItemRepository {
	return &PostgresCartItemRepository{BaseRepository: NewBaseRepository[models.CartItem](preloads)}
}

func (r *PostgresCartItemRepository) GetCartByUserId(ctx context.Context, userId int) (*models.Cart, error) {
	var cart models.Cart
	err := r.database.WithContext(ctx).Where("user_id = ?", userId).First(&cart).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
		}
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Select, err.Error())
		return nil, err
	}
	return &cart, nil
}

func (r *PostgresCartItemRepository) GetCartItemProduct(ctx context.Context, productId int) (*models.Product, error) {
	var product models.Product
	err := r.database.WithContext(ctx).Where("Id = ?", productId).First(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
		}
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Select, err.Error())
		return nil, err
	}
	return &product, nil
}
