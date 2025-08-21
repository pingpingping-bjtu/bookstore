package service

import (
	"bookstore-manager/model"
	"bookstore-manager/repository"
	"errors"
	"fmt"
	"time"
)

type OrderItems struct {
	BookID   int `json:"book_id"`
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

type CreateOrderRequest struct {
	UserID int          `json:"user_id"`
	Items  []OrderItems `json:"items"`
}

type OrderService struct {
	OrderDAO *repository.OrderDAO
	BookDao  *repository.BookDAO
}

func (o *OrderService) CreateOrder(req *CreateOrderRequest) (*model.Order, error) {
	if len(req.Items) == 0 {
		return nil, errors.New("空的订单信息")
	}
	//1.判断书籍库存是否充足、
	err := o.checkStockAvailability(req.Items)
	if err != nil {
		return nil, errors.New("书籍库存不足")
	}
	//2.生成订单号（下单成功）
	orderNo := o.generateOrderNo()
	var totalAmount int
	var orderItems []*model.OrderItem
	for _, item := range req.Items {
		subtotal := item.Price * item.Quantity
		totalAmount += subtotal
		orderItems = append(orderItems, &model.OrderItem{

			BookID:   item.BookID,
			Quantity: item.Quantity,
			Price:    item.Price,
			Subtotal: subtotal,
		})
	}

	//3.支付（判断订单下单状态）
	order := &model.Order{
		UserID:      req.UserID,
		OrderNo:     orderNo,
		TotalAmount: totalAmount,
		Status:      0,
		IsPaid:      false,
	}
	err = o.OrderDAO.CreateOrderWithItems(order, orderItems)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *OrderService) checkStockAvailability(items []OrderItems) error {
	for _, item := range items {
		book, err := o.BookDao.GetBookByID(item.BookID)
		if err != nil {
			return errors.New("图书不存在")
		}
		if book.Status != 1 {
			return errors.New("图书已下架")
		}
		if book.Stock < item.Quantity {
			return errors.New("库存不足")
		}
	}
	return nil

}

func (o *OrderService) generateOrderNo() string {
	//用时间戳去标记
	orderNo := fmt.Sprintf("ORD%d", time.Now().UnixNano())
	return orderNo
}

func NewOrderService() *OrderService {
	return &OrderService{OrderDAO: repository.NewOrderDAO(), BookDao: repository.NewBookDAO()}
}
