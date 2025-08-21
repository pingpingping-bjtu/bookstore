package controller

import (
	"bookstore-manager/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService *service.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{OrderService: service.NewOrderService()}
}

func (o *OrderController) CreateOrder(c *gin.Context) {

	var req service.CreateOrderRequest

	//绑定item参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数绑定失败",
			"error":   err.Error(),
		})
		return
	}
	//获取用户信息
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "请先登录",
		})
		return
	}
	req.UserID = userID.(int)
	order, err := o.OrderService.CreateOrder(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "订单创建失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    order,
		"message": "创建订单成功",
	})

}

func (o *OrderController) GetOrderList(c *gin.Context) {
	//获取用户信息
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "请先登录",
		})
		return
	}
}
