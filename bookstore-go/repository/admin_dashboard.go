package repository

import (
	"bookstore-manager/global"
	"bookstore-manager/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

type AdminDashboardDAO struct {
	db *gorm.DB
}

func (d *AdminDashboardDAO) GetTotalBooks() (int64, error) {
	var count int64
	if err := d.db.Debug().Model(&model.Book{}).Count(&count).Error; err != nil {
		return 0, errors.New("获取图书总量失败")
	}
	return count, nil
}

func (d *AdminDashboardDAO) GetTotalOrders() (int64, error) {
	var count int64
	if err := d.db.Debug().Model(&model.Order{}).Count(&count).Error; err != nil {
		return 0, errors.New("获取订单总量失败")
	}
	return count, nil
}

func (d *AdminDashboardDAO) GetTotalUsers() (int64, error) {
	var count int64
	if err := d.db.Debug().Model(&model.User{}).Count(&count).Error; err != nil {
		return 0, errors.New("获取用户总量失败")
	}
	return count, nil
}

func (d *AdminDashboardDAO) GetTotalRevenue() (int64, error) {
	var totalRevenue int64
	if err := d.db.Debug().Model(&model.Order{}).
		Where("is_paid = ?", true).
		Select("COALESCE(SUM(total_amount), 0)").
		Scan(&totalRevenue).Error; err != nil {
		return 0, errors.New("获取总收入失败")
	}
	return totalRevenue, nil
}

func (d *AdminDashboardDAO) GetRecentBooks() ([]model.Book, error) {
	var recentBooks []model.Book
	threeDaysAgo := time.Now().AddDate(0, 0, -3)
	err := d.db.Where("created_at >= ?", threeDaysAgo).
		Order("created_at DESC").
		Limit(10).
		Find(&recentBooks).Error
	if err != nil {
		return nil, err
	}
	return recentBooks, nil
}

func NewAdminDashboardDAO() *AdminDashboardDAO {
	return &AdminDashboardDAO{db: global.GetDB()}
}
