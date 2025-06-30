package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/domain/filter"
	"github.com/alielmi98/go-ecommerce-api/infra/db"
	"github.com/alielmi98/go-ecommerce-api/pkg/service_errors"
	"gorm.io/gorm"
)

const softDeleteExp string = "id = ? and deleted_by is null"

type BaseRepository[TEntity any] struct {
	database *gorm.DB
	preloads []db.PreloadEntity
}

func NewBaseRepository[TEntity any](preloads []db.PreloadEntity) *BaseRepository[TEntity] {
	return &BaseRepository[TEntity]{
		database: db.GetDb(),
		preloads: preloads,
	}
}

func (r BaseRepository[TEntity]) Create(ctx context.Context, entity TEntity) (TEntity, error) {
	tx := r.database.WithContext(ctx).Begin()
	err := tx.
		Create(&entity).
		Error
	if err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Insert, err.Error())
		return entity, err
	}
	tx.Commit()
	return entity, nil
}

func (r BaseRepository[TEntity]) Update(ctx context.Context, id int, entity TEntity) (TEntity, error) {
	model := new(TEntity)

	err := r.database.WithContext(ctx).Where(softDeleteExp, id).First(model).Error
	if err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Update, err.Error())
		return *model, err
	}

	*model = entity

	tx := r.database.WithContext(ctx).Begin()
	if err := tx.Model(model).Where("id = ?", id).Updates(model).Error; err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Update, err.Error())
		return *model, err
	}

	tx.Commit()
	return *model, nil
}
func (r BaseRepository[TEntity]) Delete(ctx context.Context, id int) error {
	tx := r.database.WithContext(ctx).Begin()

	model := new(TEntity)

	if ctx.Value(constants.UserIdKey) == nil {
		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}

	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"deleted_at": sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}

	if cnt := tx.
		Model(model).
		Where(softDeleteExp, id).
		Updates(deleteMap).
		RowsAffected; cnt == 0 {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Delete, service_errors.RecordNotFound)
		return &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
	}
	tx.Commit()
	return nil
}

func (r BaseRepository[TEntity]) GetById(ctx context.Context, id int) (TEntity, error) {
	model := new(TEntity)
	database := db.Preload(r.database, r.preloads)
	err := database.
		Where(softDeleteExp, id).
		First(model).
		Error
	if err != nil {
		return *model, err
	}
	return *model, nil
}

func (r BaseRepository[TEntity]) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]TEntity, error) {
	model := new(TEntity)
	var items *[]TEntity

	database := db.Preload(r.database, r.preloads)
	query := db.GenerateDynamicQuery[TEntity](&req.DynamicFilter)
	sort := db.GenerateDynamicSort[TEntity](&req.DynamicFilter)
	var totalRows int64 = 0

	database.
		Model(model).
		Where(query).
		Count(&totalRows)

	err := database.
		Where(query).
		Offset(req.GetOffset()).
		Limit(req.GetPageSize()).
		Order(sort).
		Find(&items).
		Error

	if err != nil {
		return 0, &[]TEntity{}, err
	}
	return totalRows, items, err

}

func (r *BaseRepository[TEntity]) BeginTransaction(ctx context.Context) (*gorm.DB, error) {
	tx := r.database.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}
	return tx, nil
}

func (r BaseRepository[TEntity]) CreateTx(tx *gorm.DB, entity TEntity) (TEntity, error) {
	err := tx.
		Create(&entity).
		Error
	if err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Insert, err.Error())
		return entity, err
	}
	return entity, nil
}
