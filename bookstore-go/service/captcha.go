package service

import (
	"bookstore-manager/global"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mojocn/base64Captcha"
)

type CaptchaService struct {
	store base64Captcha.Store
}

func NewCaptchaService() *CaptchaService {
	return &CaptchaService{
		store: base64Captcha.DefaultMemStore,
	}
}

type CaptchaResponse struct {
	CaptchaID     string `json:"captcha_id"`
	CaptchaBase64 string `json:"captcha_base_64"`
}

func (c *CaptchaService) GenerateCaptcha() (*CaptchaResponse, error) {
	//验证码格式
	driver := base64Captcha.NewDriverDigit(
		80,
		240,
		4,   //验证码长度
		0.7, //干扰强度
		80)
	captcha := base64Captcha.NewCaptcha(driver, c.store)
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		return nil, err
	}

	//用redis,作用是存储有效期的图片验证码
	ctx := context.Background()
	log.Println("图片验证码真实answer:", answer)
	redisKey := fmt.Sprintf("captcha:%s", id)
	//key是redisKey ，value answer
	err = global.RedisClient.Set(ctx, redisKey, answer, 1*time.Minute).Err()
	if err != nil {
		return nil, err
	}
	return &CaptchaResponse{
		CaptchaID:     id,
		CaptchaBase64: b64s,
	}, nil
}

// VerifyCaptcha 验证验证码
func (c *CaptchaService) VerifyCaptcha(captchaID, captchaValue string) bool {
	if captchaValue == "" || captchaID == "" {
		return false
	}
	//从redis获取验证码答案answer
	ctx := context.Background()
	redisKey := fmt.Sprintf("captcha:%s", captchaID)
	storedAnswer, err := global.RedisClient.Get(ctx, redisKey).Result()
	if err != nil {
		return false
	}
	//比较用户输入的验证码和存储的答案
	isValid := storedAnswer == captchaValue
	//验证成功后删除redis的验证码
	if isValid {
		global.RedisClient.Del(ctx, redisKey)
	}
	return isValid
}
