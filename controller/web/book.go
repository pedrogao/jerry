package web

import (
	Err "github.com/PedroGao/jerry/libs/err"
	"github.com/PedroGao/jerry/model"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/builder"
	"net/http"
	"strconv"
	"strings"
)

func GetBook(c *gin.Context) {
	var (
		err error
		i   int
	)
	id := c.Param("id")
	i, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(Err.ParamsErr.HTTPCode, Err.ParamsErr.SetUrl(c.Request.URL.String()).Clear().Add("参数错误，id必须为正整数"))
		return
	}
	book, err := model.GetBookById(i)
	if err != nil {
		c.JSON(Err.BookNotFound.HTTPCode, Err.BookNotFound.SetUrl(c.Request.URL.String()))
		return
	}
	c.JSON(http.StatusOK, book)
}

func SearchBook(c *gin.Context) {
	var (
		err   error
		books []model.Book
	)
	keyword := c.Query("keyword")
	if strings.TrimSpace(keyword) == "" {
		c.JSON(Err.ParamsErr.HTTPCode, Err.ParamsErr.SetUrl(c.Request.URL.String()).Clear().Add("参数错误，关键字必须有效词"))
		return
	}
	// limit 5 for test
	err = model.DB.Where(builder.Like{"title", keyword}).Limit(5).Find(&books)
	if err != nil {
		c.JSON(Err.ParamsErr.HTTPCode, Err.ParamsErr.SetUrl(c.Request.URL.String()).Clear().Add(err.Error()))
		return
	}
	if len(books) < 1 {
		c.JSON(Err.BookNotFound.HTTPCode, Err.BookNotFound.SetUrl(c.Request.URL.String()))
		return
	}
	c.JSON(http.StatusOK, books)
}
