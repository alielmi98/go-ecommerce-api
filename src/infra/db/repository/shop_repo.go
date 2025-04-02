package repository

import (
	"context"
	"log"

	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/infra/db"
	"gorm.io/gorm"
)

const softDeleteExp string = "id = ? and deleted_by is null"

type BaseRepository[TEntity any] struct {
	database *gorm.DB
}

func NewBaseRepository[TEntity any]() *BaseRepository[TEntity] {
	return &BaseRepository[TEntity]{
		database: db.GetDb(),
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
