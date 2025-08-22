package repository

import (
	"bookstore-manager/global"
	"bookstore-manager/model"
	"errors"

	"gorm.io/gorm"
)

type AdminUserDAO struct {
	db *gorm.DB
}

func (u *AdminUserDAO) GetUserByUsername(username string) (*model.User, error) {
	var user *model.User
	err := u.db.Debug().Where("username=?", username).Find(&user).Error
	if err != nil {
		return nil, errors.New("用户名不正确")
	}
	return user, nil
}

func NewAdminUserDAO() *AdminUserDAO {
	return &AdminUserDAO{db: global.GetDB()}
}
