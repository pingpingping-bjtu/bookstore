package controller

import (
	"bookstore-manager/service"
	"net/http"
	"strconv"

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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "12"))
	//获取用户信息
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "请先登录",
		})
		return
	}
	orders, total, err := o.OrderService.GetUserOrders(userID.(int), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "获取订单列表失败",
			"error":   err,
		})
		return
	}
	totalPages := total + (int64(pageSize-1))/int64(pageSize)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取订单列表成功",
		"data": gin.H{
			"orders":      orders,
			"total":       total,
			"page":        page,
			"page_size":   pageSize,
			"total_pages": totalPages,
		},
	})
}

func (o *OrderController) PayOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "无效的订单号",
		})
		return
	}
	err = o.OrderService.PayOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "支付失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "支付成功",
	})

}
