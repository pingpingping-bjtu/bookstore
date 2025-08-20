package repository

import (
	"bookstore-manager/global"
	"bookstore-manager/model"

	"gorm.io/gorm"
)

type BookDAO struct {
	db *gorm.DB
}

func NewBookDAO() *BookDAO {
	return &BookDAO{
		db: global.GetDB(),
	}
}
func (b *BookDAO) GetHotBooks(limit int) ([]*model.Book, error) {
	var books []*model.Book
	err := b.db.Debug().Where("status=?", 1).Order("sale desc").Limit(limit).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (b *BookDAO) GetNewBooks(limit int) ([]*model.Book, error) {
	var books []*model.Book
	err := b.db.Debug().Where("status=?", 1).Order("created_at desc").Limit(limit).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b *BookDAO) GetBookByPage(page int, pageSize int) ([]*model.Book, int64, error) {
	var books []*model.Book
	var total int64
	err := b.db.Model(&model.Book{}).Debug().Where("status=?", 1).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	//利用sql的offset语法实现位移
	offset := (page - 1) * pageSize
	err = b.db.Debug().Where("status=?", 1).Offset(offset).Limit(pageSize).Find(&books).Error
	if err != nil {
		return nil, total, err
	}
	return books, total, nil
}

func (b *BookDAO) SearchBooksWithPage(keyword string, page int, pageSize int) ([]*model.Book, int64, error) {
	var books []*model.Book
	var total int64
	searchCondition := b.db.Debug().Where("status=? AND (title LIKE ? OR author LIKE ? OR description LIKE ?)",
		1, "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	//记录搜索出的书籍总数
	err := searchCondition.Model(&model.Book{}).Debug().Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	//分页逻辑
	offSet := (page - 1) * pageSize
	err = searchCondition.Offset(offSet).Limit(pageSize).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}
	return books, total, nil
}

func (b *BookDAO) GetBookByID(id int) (*model.Book, error) {
	var book *model.Book
	err := b.db.Debug().Where("status=?", 1).First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}
