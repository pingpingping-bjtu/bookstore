package service

import (
	"bookstore-manager/jwt"
	"bookstore-manager/model"
	"bookstore-manager/repository"
	"encoding/base64"

	"errors"
)

type UserService struct {
	UserDB *repository.UserDAO
}

type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpireIn     int64     `json:"expire_in"`
	UserInfo     *UserInfo `json:"user_info"`
}
type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// NewUserService service => repository => 调用db方法（对应 model 的模型）
func NewUserService() *UserService {
	return &UserService{
		UserDB: repository.NewUserDAO(),
	}
}

// UserRegister 用户信息注册
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

// UserLogin 用户登录校验
func (u *UserService) UserLogin(username, password string) (*LoginResponse, error) {
	//1.查询用户是否存在
	user, err := u.UserDB.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("用户不存在！")
	}
	//2.如果存在，验证密码是否正确
	if !u.VersifyPassword(password, user.Password) {
		return nil, errors.New("密码错误")
	}
	//3.JWT
	token, err := jwt.GenerateTokenPair(uint(user.ID), user.Username)
	if err != nil {
		return nil, errors.New("生成token失败")
	}
	response := &LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpireIn:     token.ExpiresIn,
		UserInfo: &UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		},
	}
	return response, nil
}

func (u *UserService) VersifyPassword(inputPassword, TruePassword string) bool {
	return u.encodePassword(inputPassword) == TruePassword
}

func (u *UserService) GetUserByID(userID int) (*model.User, error) {
	user, err := u.UserDB.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (u *UserService) UpdateUserInfo(user *model.User) error {
	//看这个用户在不在
	oldUser, err := u.UserDB.GetUserByID(user.ID)
	if err != nil {
		return err
	}
	oldUser.Phone = user.Phone
	oldUser.Email = user.Email
	oldUser.Avatar = user.Avatar
	oldUser.Username = user.Username
	//调用dao层更新信息
	err = u.UserDB.UpdateUser(oldUser)
	if err != nil {
		return err
	}
	return err
}
