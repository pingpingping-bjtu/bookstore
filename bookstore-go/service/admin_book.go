package service

import (
	"bookstore-manager/model"
	"bookstore-manager/repository"
)

type AdminBookService struct {
	AdminBookDAO *repository.AdminBookDAO
}

func (b *AdminBookService) GetBookList(page, pageSize int, req *repository.GetBookRequest) (*[]model.Book, int64, error) {
	return b.AdminBookDAO.GetBookList(page, pageSize, req)

}

func (b *AdminBookService) CreateBook(req *model.BookCreateRequest) error {
	// 确保CategoryID有有效值
	categoryID := req.CategoryID
	if categoryID <= 0 {
		categoryID = 1 // 默认使用第一个分类
	}

	book := &model.Book{
		Title:       req.Title,
		Author:      req.Author,
		Price:       req.Price,
		Discount:    req.Discount,
		Type:        req.Type,
		Stock:       req.Stock,
		Status:      req.Status,
		CoverURL:    req.CoverURL,
		Description: req.Description,
		ISBN:        req.ISBN,
		Publisher:   req.Publisher,
		PublishDate: req.PublishDate,
		Pages:       req.Pages,
		Language:    req.Language,
		Format:      req.Format,
		CategoryID:  categoryID,
		Sale:        req.Sale,
	}

	return b.AdminBookDAO.CreateBook(book)
}

func (b *AdminBookService) GetBookByID(id int) (*model.Book, error) {
	return b.AdminBookDAO.GetBookByID(id)
}

func (b *AdminBookService) UpdateBookStatus(id uint, status int) error {
	book, err := b.AdminBookDAO.GetBookByID(int(id))
	if err != nil {
		return err
	}
	book.Status = status
	return b.AdminBookDAO.UpdateBookStatus(book)
}

func (b *AdminBookService) DeleteBook(id uint64) error {
	return b.AdminBookDAO.DeleteBook(int(id))
}

func (b *AdminBookService) UpdateBook(id int, req *model.BookUpdateRequest) error {
	book, err := b.GetBookByID(id)
	if err != nil {
		return err
	}
	// 更新字段（排除状态字段，避免意外修改状态）
	if req.Title != "" {
		book.Title = req.Title
	}
	if req.Author != "" {
		book.Author = req.Author
	}
	if req.Price > 0 {
		book.Price = req.Price
	}
	if req.Discount > 0 {
		book.Discount = req.Discount
	}
	if req.Type != "" {
		book.Type = req.Type
	}
	if req.Stock >= 0 {
		book.Stock = req.Stock
	}
	// 注意：不更新 Status 字段，避免意外修改状态
	// if req.Status >= 0 {
	// 	book.Status = req.Status
	// }
	if req.CoverURL != "" {
		book.CoverURL = req.CoverURL
	}
	if req.Description != "" {
		book.Description = req.Description
	}
	if req.ISBN != "" {
		book.ISBN = req.ISBN
	}
	if req.Publisher != "" {
		book.Publisher = req.Publisher
	}
	if req.PublishDate != "" {
		book.PublishDate = req.PublishDate
	}
	if req.Pages > 0 {
		book.Pages = req.Pages
	}
	if req.Language != "" {
		book.Language = req.Language
	}
	if req.Format != "" {
		book.Format = req.Format
	}
	if req.CategoryID > 0 {
		book.CategoryID = req.CategoryID
	}
	if req.Sale >= 0 {
		book.Sale = req.Sale
	}

	return b.AdminBookDAO.UpdateBook(book)
}

func NewAdminBookService() *AdminBookService {
	return &AdminBookService{AdminBookDAO: repository.NewAdminBookDAO()}
}
