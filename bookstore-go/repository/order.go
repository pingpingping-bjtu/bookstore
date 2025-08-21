package repository

import (
	"bookstore-manager/global"
	"bookstore-manager/model"

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

func NewOrderDAO() *OrderDAO {
	return &OrderDAO{db: global.GetDB()}
}
