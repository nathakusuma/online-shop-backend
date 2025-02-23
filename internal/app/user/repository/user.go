package repository

import (
	"gorm.io/gorm"
	"online-shop-backend/internal/domain/entity"
)

type UserMySQLItf interface {
	CreateUser(user entity.User) error
	GetUserByEmail(email string) (entity.User, error)
}

type UserMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) UserMySQLItf {
	return &UserMySQL{db}
}

func (u *UserMySQL) CreateUser(user entity.User) error {
	return u.db.Create(&user).Error
}

func (u *UserMySQL) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return user, err
}
