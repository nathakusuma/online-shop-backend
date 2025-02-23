package mysql

import (
	"gorm.io/gorm"
	"online-shop-backend/internal/domain/entity"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		entity.User{},
		entity.Product{},
	)
}
