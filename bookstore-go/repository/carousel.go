package repository

import (
	"bookstore-manager/global"
	"bookstore-manager/model"

	"gorm.io/gorm"
)

type CarouselDAO struct {
	db *gorm.DB
}

func (c *CarouselDAO) GetCarouselList() ([]*model.Carousel, error) {
	var carousels []*model.Carousel
	err := c.db.Debug().Find(&carousels).Error
	if err != nil {
		return nil, err
	}
	return carousels, nil
}

func NewCarouselDAO() *CarouselDAO {
	return &CarouselDAO{db: global.GetDB()}
}
