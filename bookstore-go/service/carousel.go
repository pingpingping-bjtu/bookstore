package service

import (
	"bookstore-manager/model"
	"bookstore-manager/repository"
)

type CarouselService struct {
	CarouselDAO *repository.CarouselDAO
}

func (c *CarouselService) GetCarouselList() ([]*model.Carousel, error) {
	return c.CarouselDAO.GetCarouselList()
}

func NewCarouselService() *CarouselService {
	return &CarouselService{CarouselDAO: repository.NewCarouselDAO()}
}
