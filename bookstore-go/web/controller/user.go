package controller

import (
	"bookstore-manager/jwt"
	"bookstore-manager/model"
	"bookstore-manager/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	CaptchaID       string `json:"captcha_id"`
	CaptchaValue    string `json:"captcha_value"`
}

type LoginRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaID    string `json:"captcha_id"`
	CaptchaValue string `json:"captcha_value"`
}

func (u *UserController) UserRegister(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "参数绑定失败",
			"error":   err,
		})
	}
	//验证码校验
	captchaSvc := service.NewCaptchaService()
	if !captchaSvc.VerifyCaptcha(req.CaptchaID, req.CaptchaValue) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "验证码错误",
		})
		return
	}
	//验证密码两次是否一致
	if req.Password != req.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "两次密码不一致",
		})
		return
	}
	if err := u.UserService.UserRegister(req.Username, req.Password, req.Phone, req.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "注册成功",
	})

}

func (u *UserController) UserLogin(ctx *gin.Context) {
	//JWT 用于验证用户身份的哈希值，并且服务端针对哈希值获取相应的用户信息
	//1.验证图片验证码
	//2.校验用户信息，是否有这个用户
	//3.返回JWT信息给用户，后面发送信息就知道是哪个用户
	var req LoginRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
	}
	captchaSvc := service.NewCaptchaService()
	if !captchaSvc.VerifyCaptcha(req.CaptchaID, req.CaptchaValue) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "验证码错误",
		})
		return
	}
	response, err := u.UserService.UserLogin(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    response,
		"message": "登录成功",
	})
}

func (u *UserController) GetUserProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "用户未登录",
		})
	}
	//调用服务层获取用户信息
	user, err := u.UserService.GetUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	response := gin.H{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"phone":      user.Phone,
		"avatar":     user.Avatar,
		"created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at": user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    response,
		"message": "获取用户信息成功",
	})
}

func (u *UserController) UpdateUserProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "用户未登录",
		})
	}
	var updateData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
	}
	if err := c.ShouldBindBodyWithJSON(&updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "参数绑定错误",
			"error":   err.Error(),
		})
		return
	}
	user := &model.User{
		ID:       userID.(int),
		Username: updateData.Username,
		Email:    updateData.Email,
		Phone:    updateData.Phone,
		Avatar:   updateData.Avatar,
	}
	err := u.UserService.UpdateUserInfo(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "用户信息更新失败",
			"error":   err.Error(),
		})
	}
	//获取更新信息
	updateUser, err := u.UserService.GetUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "用户信息更新失败",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    updateUser,
		"message": "更新用户信息成功",
	})

}

func (u *UserController) ChangePassword(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "用户未登录",
		})
	}
	var passwordData struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindBodyWithJSON(&passwordData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "参数绑定错误",
			"error":   err.Error(),
		})
		return
	}
	//前端已经添加长度限制规则，后端不需要
	//if len(passwordData.NewPassword) < 6 {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code":    -1,
	//		"message": "新密码至少六位",
	//	})
	//
	//	return
	//}
	err := u.UserService.ChangePassword(userID.(int), passwordData.OldPassword, passwordData.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    passwordData,
		"message": "密码更新成功",
	})
}

func (u *UserController) LogOut(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "用户未登录",
		})
	}
	//撤销用户的token
	err := jwt.RevokeToken(uint(userID.(int)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "退出登录失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "退出成功",
	})
}
