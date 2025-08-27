package repository

import (
	"bookstore-manager/global"
	"bookstore-manager/model"
	"strconv"

	"gorm.io/gorm"
)

type AdminBookDAO struct {
	db *gorm.DB
}
type GetBookRequest struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Status   string `json:"status"`
	BookType string `json:"type"`
}

func (b *AdminBookDAO) GetBookList(page, pageSize int, req *GetBookRequest) (*[]model.Book, int64, error) {
	var books *[]model.Book
	var total int64
	query := b.db.Debug().Model(&model.Book{})
	query.Count(&total)
	offset := (page - 1) * pageSize
	if req.Author != "" {
		query = query.Where("author LIKE ?", "%"+req.Author+"%")
	}
	if req.Title != "" {
		query = query.Where("title LIKE ?", "%"+req.Title+"%")
	}
	if req.BookType != "" {
		query = query.Where("type=?", req.BookType)
	}
	if req.Status != "" {
		status, _ := strconv.Atoi(req.Status)
		query = query.Where("status=?", status)
	}
	err := query.Offset(offset).Limit(pageSize).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}
	return books, total, nil
}

func (b *AdminBookDAO) CreateBook(book *model.Book) error {
	return b.db.Debug().Create(&book).Error
}

func (b *AdminBookDAO) GetBookByID(id int) (*model.Book, error) {
	var book *model.Book
	err := b.db.Debug().Where("id=?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (b *AdminBookDAO) UpdateBookStatus(book *model.Book) error {
	return b.db.Debug().Save(&book).Error

}

func (b *AdminBookDAO) DeleteBook(id int) error {
	return b.db.Debug().Delete(&model.Book{}, id).Error
}

func (b *AdminBookDAO) UpdateBook(book *model.Book) error {
	return b.db.Debug().Save(book).Error
}

func NewAdminBookDAO() *AdminBookDAO {
	return &AdminBookDAO{db: global.GetDB()}
}
