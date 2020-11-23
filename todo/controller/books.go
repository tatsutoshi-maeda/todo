package controller

import (
	"database/sql"
	"net/http"
	"strconv"
	"todo/model"
	"todo/models"
	"todo/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func BookAdd(ctx *gin.Context) {
	DB, _ := sql.Open("mysql", "test_u:test_pw@tcp(mysql:3306)/todo_db?charset=utf8")
	book := models.Book{}
	err := ctx.Bind(&book)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	err = book.Insert(ctx, DB, boil.Infer())
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Server Error")
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookList(ctx *gin.Context) {
	bookService := service.BookService{}
	BookLists := bookService.GetBookList()
	ctx.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    BookLists,
	})
}

// func BookList(ctx *gin.Context) {
// 	DB, _ := sql.Open("mysql", "test_u:test_pw@(127.0.0.1:3307)/todo_db?charset=utf8")
// 	book := models.Book{}
// 	BookLists := book.All(ctx, DB)
// 	ctx.JSONP(http.StatusOK, gin.H{
// 		"message": "ok",
// 		"data":    BookLists,
// 	})
// }

func BookUpdate(ctx *gin.Context) {
	book := model.Book{}
	err := ctx.Bind(&book)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	err = bookService.UpdateBook(&book)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Server Error")
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookDelete(ctx *gin.Context) {
	id := ctx.PostForm("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	err = bookService.DeleteBook(int(intId))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Server Error")
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
