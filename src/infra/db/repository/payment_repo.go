package repository

import (
	"context"

	"github.com/alielmi98/go-ecommerce-api/domain/models"
	"github.com/alielmi98/go-ecommerce-api/infra/db"
	"github.com/alielmi98/go-ecommerce-api/pkg/service_errors"
	"gorm.io/gorm"
)

type PostgresPaymentRepository struct {
	*BaseRepository[models.Payment]
}

func NewPaymentRepository(preloads []db.PreloadEntity) *PostgresPaymentRepository {
	return &PostgresPaymentRepository{BaseRepository: NewBaseRepository[models.Payment](preloads)}
}

func (r *PostgresPaymentRepository) GetByAuthority(ctx context.Context, authority string) (*models.Payment, error) {
	var payment models.Payment
	err := r.database.WithContext(ctx).Where("authority_id = ?", authority).First(&payment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
		}
		return nil, err
	}
	return &payment, nil
}
