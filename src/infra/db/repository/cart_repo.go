package repository

import (
	"context"
	"log"

	"github.com/alielmi98/go-ecommerce-api/constants"

	"github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/infra/db"
	"github.com/alielmi98/go-ecommerce-api/pkg/service_errors"
	"gorm.io/gorm"
)

type PostgresCartRepository struct {
	*BaseRepository[models.Cart]
}

func NewCartRepository(preloads []db.PreloadEntity) *PostgresCartRepository {
	return &PostgresCartRepository{BaseRepository: NewBaseRepository[models.Cart](preloads)}
}

func (r *PostgresCartRepository) GetCartByUserId(ctx context.Context, userId int) (*models.Cart, error) {
	var cart models.Cart
	database := db.Preload(r.database, r.preloads)
	err := database.WithContext(ctx).Where("user_id = ?", userId).First(&cart).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
		}
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Select, err.Error())
		return nil, err
	}
	return &cart, nil
}
