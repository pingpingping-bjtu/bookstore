package service

import (
	"bookstore-manager/global"
	"bookstore-manager/jwt"
	"bookstore-manager/model"
	"bookstore-manager/repository"

	"encoding/base64"
	"errors"
)

type AdminLoginResponse struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

type AdminUserService struct {
	AdminUserDAO *repository.AdminUserDAO
}

func (u *AdminUserService) AdminUserLogin(username, password string) (*AdminLoginResponse, error) {
	//获取user，不存在返回用户名错误
	user, err := u.AdminUserDAO.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	//验证密码是否正确
	if !u.VerifyPassword(password, user.Password) {
		return nil, errors.New("密码不正确")
	}
	//验证是否是管理员
	if !user.IsAdmin {
		return nil, errors.New("非管理员用户")
	}
	//返回token
	token, err := jwt.GenerateToken(uint(user.ID), user.Username)
	if err != nil {
		return nil, errors.New("生成token失败")
	}
	response := &AdminLoginResponse{
		Token: token,
		User:  user,
	}
	return response, nil
}

func NewAdminUserService() *AdminUserService {
	return &AdminUserService{AdminUserDAO: repository.NewAdminUserDAO()}
}

// VerifyPassword 两个参数都不需要编码
func (u *AdminUserService) VerifyPassword(inputPassword, TruePassword string) bool {
	return u.encodePassword(inputPassword) == TruePassword
}
func (u *AdminUserService) encodePassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}

func (u *AdminUserService) GetUsersList(req *repository.GetUsersRequest, page int, size int) (*[]model.User, int64, error) {
	return u.AdminUserDAO.GetUsersList(req, page, size)

}

func (u *AdminUserService) CreateUser(req *repository.CreateUserRequest) (*model.User, error) {
	//1.检查用户名，邮箱，校验和唯一性
	flag, err := u.AdminUserDAO.CheckUserExists(req.Username, req.Phone, req.Email)
	if err != nil {
		return nil, err
	}
	if flag {
		return nil, errors.New("用户名、手机号、邮箱已存在")
	}
	//密码加密
	//2.密码加密
	encodePassword := u.encodePassword(req.Password)

	user, err := u.AdminUserDAO.CreateUser(req.Username, encodePassword, req.Phone, req.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *AdminUserService) DeleteUser(id uint) error {
	user, err := u.getUserByID(id)
	if err != nil {
		return errors.New("用户不存在")
	}
	return u.AdminUserDAO.DeleteUser(user)

}

func (u *AdminUserService) getUserByID(id uint) (*model.User, error) {
	return u.AdminUserDAO.GetUserByID(id)
}

func (u *AdminUserService) UpdateUserStatus(isAdmin bool, id uint) error {
	return u.AdminUserDAO.UpdateUserStatus(isAdmin, id)
}

func (u *AdminUserService) UpdateUser(id uint, req *repository.UpdateUserRequest) (*model.User, error) {
	//检查用户是否存在
	_, err := u.getUserByID(id)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	// 更新字段
	updates := make(map[string]interface{})
	if req.Username != "" {
		// 检查用户名是否已被其他用户使用
		var existingUser model.User
		if err := global.DBClient.Where("username = ? AND id != ?", req.Username, id).First(&existingUser).Error; err == nil {
			return nil, errors.New("用户名已存在")
		}
		updates["username"] = req.Username
	}

	if req.Email != "" {
		// 检查邮箱是否已被其他用户使用
		var existingUser model.User
		if err := global.DBClient.Where("email = ? AND id != ?", req.Email, id).First(&existingUser).Error; err == nil {

			return nil, errors.New("邮箱已被使用")
		}
		updates["email"] = req.Email
	}

	if req.Phone != "" {
		updates["phone"] = req.Phone
	}

	if req.IsAdmin != nil {
		updates["is_admin"] = *req.IsAdmin
	}

	return u.AdminUserDAO.UpdateUserService(id, updates)
}
