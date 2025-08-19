package repository

import (
	"bookstore-manager/global"
	"bookstore-manager/model"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO() *UserDAO {
	return &UserDAO{
		db: global.GetDB(),
	}
}

func (u *UserDAO) CreateUser(user *model.User) error {
	if err := u.db.Debug().Create(user).Error; err != nil {
		return err
	}
	return nil
}

// CheckUserExists 检查user是否存在
func (u *UserDAO) CheckUserExists(username, phone, email string) (bool, error) {
	var count int64
	err := u.db.Model(&model.User{}).Where("username=?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, err
	}

	err = u.db.Model(&model.User{}).Where("phone=?", phone).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, err
	}

	err = u.db.Model(&model.User{}).Where("email=?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, err
	}
	return count > 0, nil
}
