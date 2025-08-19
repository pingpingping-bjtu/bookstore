package controller

import (
	"bookstore-manager/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CaptchaController struct {
	CaptchaService *service.CaptchaService
}

func NewCaptchaController() *CaptchaController {
	return &CaptchaController{CaptchaService: service.NewCaptchaService()}
}

// GenerateCaptcha   生成图形验证码
func (c *CaptchaController) GenerateCaptcha(ctx *gin.Context) {
	//生成图形验证码
	res, err := c.CaptchaService.GenerateCaptcha()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "生成验证码失败",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    res,
		"message": "生成验证码成功",
	})

}
