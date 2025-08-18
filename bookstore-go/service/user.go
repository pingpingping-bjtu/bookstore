package service

import (
	"bookstore-manager/model"
	"bookstore-manager/repository"
	"encoding/base64"

	"errors"
)

type UserService struct {
	UserDB *repository.UserDAO
}

// NewUserService service => repository => 调用db方法（对应 model 的模型）
func NewUserService() *UserService {
	return &UserService{
		UserDB: repository.NewUserDAO(),
	}
}

// UserRegister 服务注册
func (u *UserService) UserRegister(username, password, phone, email string) error {
	//1.检查用户名，邮箱，校验和唯一性
	flag, err := u.UserDB.CheckUserExists(username, password, email)
	if err != nil {
		return err
	}
	if flag {
		return errors.New("用户名、手机号、邮箱已存在")
	}
	//2.密码加密
	encodepassword := u.encodePassword(password)

	err = u.createUser(username, encodepassword, phone, email)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserService) encodePassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}

func (u *UserService) createUser(username, passwordHash, phone, email string) error {
	user := &model.User{
		Username: username,
		Password: passwordHash,
		Phone:    phone,
		Email:    email,
	}
	return u.UserDB.CreateUser(user)
}
