package repository

import (
	"bookstore-manager/global"
	"bookstore-manager/model"
	"errors"

	"gorm.io/gorm"
)

type OrderDAO struct {
	db *gorm.DB
}

func (o *OrderDAO) CreateOrderWithItems(order *model.Order, items []*model.OrderItem) error {
	//事务
	err := o.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Debug().Error; err != nil {
			return err
		}
		//创建订单项
		for _, item := range items {
			item.OrderID = order.ID
			if err := tx.Create(item).Debug().Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderDAO) GetUserOrders(userID int, page int, pageSize int) ([]*model.Order, int64, error) {
	var orders []*model.Order
	var total int64
	err := o.db.Debug().Model(&model.Order{}).Where("user_id=?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	//分页
	offset := (page - 1) * pageSize
	err = o.db.Preload("OrderItems.Book").Where("user_id=?", userID).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&orders).Error
	if err != nil {
		return nil, 0, err
	}
	return orders, total, nil

}

func (o *OrderDAO) UpdateOrderStatus(order *model.Order, orderID int) error {
	//是否需要更新多张表
	//订单号，销量的更新，库存的减少，金额的更新，订单的状态（0或1）
	//用事务（上面的事务当成整体）
	err := o.db.Debug().Transaction(func(tx *gorm.DB) error {
		//并发场景
		//两次检查库存
		for _, item := range order.OrderItems {
			var book *model.Book
			if err := tx.First(&book, item.BookID).Error; err != nil {
				return errors.New("图书不存在")
			}
			if book.Stock < item.Quantity {
				return errors.New("图书库存不足")
			}
		}
		if err := tx.Model(&model.Order{}).Where("id=?", orderID).Updates(
			map[string]interface{}{
				"status":       1,
				"is_paid":      true,
				"payment_time": gorm.Expr("NOW()"),
			}).Error; err != nil {
			return err
		}

		//进行销量和库存的更新
		for _, item := range order.OrderItems {
			if err := tx.Model(&model.Book{}).Where("id=?", item.BookID).Updates(
				map[string]interface{}{
					"stock": gorm.Expr("stock-?", item.Quantity),
					"sale":  gorm.Expr("sale+?", item.Quantity),
				}).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderDAO) GetOrderByID(orderID int) (*model.Order, error) {
	var order *model.Order
	err := o.db.Debug().Preload("OrderItems.Book").Find(&order, orderID).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func NewOrderDAO() *OrderDAO {
	return &OrderDAO{db: global.GetDB()}
}
