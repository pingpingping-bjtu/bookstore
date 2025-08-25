package model

import "time"

type Book struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`       //书名
	Author      string    `json:"author"`      //作者
	Price       int       `json:"price"`       //价格
	Discount    string    `json:"discount"`    //折扣
	Type        string    `json:"type"`        //类型
	Stock       int       `json:"stock"`       //库存
	Status      int       `json:"status"`      //上架1 下架0
	Description string    `json:"description"` //描述
	CoverURL    string    `json:"cover_url"`   //封面
	ISBN        string    `json:"isbn"`
	Publisher   string    `json:"publisher"` //出版社
	Pages       int       `json:"pages"`     //页数
	Language    string    `json:"language"`  //语言
	Format      string    `json:"format"`    //装订本
	CategoryID  string    `json:"category_id"`
	Sale        int       `json:"sale"` //销量
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (b *Book) TableName() string {
	return "books"
}
