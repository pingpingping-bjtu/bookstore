package service

import (
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
