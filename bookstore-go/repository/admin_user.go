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
type GetUsersRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	IsAdmin  string `json:"is_admin"`
}
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IsAdmin  *bool  `json:"is_admin"`
}

func (u *AdminUserDAO) GetUserByUsername(username string) (*model.User, error) {
	var user *model.User
	err := u.db.Debug().Where("username=?", username).Find(&user).Error
	if err != nil {
		return nil, errors.New("用户名不正确")
	}
	return user, nil
}

func (u *AdminUserDAO) GetUsersList(req *GetUsersRequest, page int, size int) (*[]model.User, int64, error) {
	var users *[]model.User
	var total int64
	query := u.db.Debug().Model(&model.User{})
	if req.Username != "" {
		query = query.Where("username LiKE ?", "%"+req.Username+"%")
	}
	if req.Email != "" {
		query = query.Where("email LiKE ?", "%"+req.Email+"%")
	}
	if req.IsAdmin != "" {
		isAdminBool := req.IsAdmin == "true"
		query = query.Where("is_admin= ?", isAdminBool)
	}
	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	//分页查询
	offset := (page - 1) * size
	err := query.Offset(offset).Limit(size).Order("created_at DESC").Find(&users).Error
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (u *AdminUserDAO) CheckUserExists(username, phone, email string) (bool, error) {
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

func (u *AdminUserDAO) CreateUser(username, password, phone, email string) (*model.User, error) {
	user := &model.User{
		Phone:    phone,
		Username: username,
		Password: password,
		Email:    email,
	}
	return user, u.db.Create(&user).Debug().Error

}

func (u *AdminUserDAO) GetUserByID(id uint) (*model.User, error) {
	var user *model.User
	err := u.db.Debug().Find(&user, id).Error
	return user, err
}

func (u *AdminUserDAO) DeleteUser(user *model.User) error {
	return u.db.Debug().Delete(&user).Error
}

func (u *AdminUserDAO) UpdateUserStatus(isAdmin bool, id uint) error {
	return u.db.Debug().Model(&model.User{}).Where("id=?", id).Update("is_admin", isAdmin).Error

}

func (u *AdminUserDAO) UpdateUserService(id uint, updates map[string]interface{}) (*model.User, error) {
	var user *model.User
	err := u.db.Debug().Model(&user).Where("id=?", id).Updates(updates).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func NewAdminUserDAO() *AdminUserDAO {
	return &AdminUserDAO{db: global.GetDB()}
}
