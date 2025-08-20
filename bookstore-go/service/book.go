package service

import (
	"bookstore-manager/model"
	"bookstore-manager/repository"
)

type BookService struct {
	BookDB *repository.BookDAO
	//TODO:类别的DAO
}

func NewBookService() *BookService {
	return &BookService{
		BookDB: repository.NewBookDAO(),
	}
}

func (b *BookService) GetHotBooks(limit int) ([]*model.Book, error) {
	return b.BookDB.GetHotBooks(limit)
}

func (b *BookService) GetNewBooks(limit int) ([]*model.Book, error) {
	return b.BookDB.GetNewBooks(limit)
}

func (b *BookService) GetBookByPage(page int, pageSize int) ([]*model.Book, int64, error) {
	return b.BookDB.GetBookByPage(page, pageSize)
}
