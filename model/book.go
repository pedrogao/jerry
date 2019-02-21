package model

import (
	"errors"
	"github.com/go-xorm/builder"
	"time"
)

type Book struct {
	CreateTime time.Time `xorm:"TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `xorm:"TIMESTAMP" json:"update_time"`
	DeleteTime time.Time `xorm:"TIMESTAMP" json:"delete_time"`
	Id         int       `xorm:"not null pk autoincr INT(11)" json:"id"`
	Title      string    `xorm:"not null VARCHAR(50)" json:"title"`
	Author     string    `xorm:"VARCHAR(30)" json:"author"`
	Summary    string    `xorm:"VARCHAR(1000)" json:"summary"`
	Image      string    `xorm:"VARCHAR(50)" json:"image"`
}

// add db query here
func GetBookById(id int) (*Book, error) {
	var book = &Book{Id: id}
	ok, err := DB.Get(book)
	if !ok {
		return nil, err
	}
	return book, nil
}

func SearchBookByKeyword(keyword string) (*Book, error) {
	book := new(Book)
	ok, _ := DB.Where(builder.Like{"title", keyword}).Get(book)
	if !ok {
		return nil, errors.New("未找到任何相关书籍")
	}
	// xorm 在找不到book时不会报错而是会返回一个空结果，所以此处还需判断book结果是否为空
	//if book == nil {
	//	return nil, errors.New("未找到任何相关书籍")
	//}
	return book, nil
}
