package controller

import (
	"bookstore-manager/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserRegister(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "参数绑定失败",
			"error":   err,
		})
	}
	//TODO:验证码的校验
	svc := service.NewUserService()
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
	if err := svc.UserRegister(req.Username, req.Password, req.Phone, req.Email); err != nil {
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

func UserLogin(ctx *gin.Context) {

}
