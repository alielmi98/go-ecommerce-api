package db

import (
	"gorm.io/gorm"
)

type PreloadEntity struct {
	Entity string
}

// Preload
func Preload(db *gorm.DB, preloads []PreloadEntity) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.Entity)
	}
	return db
}
