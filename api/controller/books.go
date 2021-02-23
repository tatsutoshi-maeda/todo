package controller

import (
	"api/models"
	"database/sql"
	"net/http"

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
	DB, _ := sql.Open("mysql", "test_u:test_pw@tcp(mysql:3306)/todo_db?charset=utf8")
	var err error
	books := []*models.Book{}

	books, err = models.Books().All(ctx, DB)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Server Error")
		return
	}
	ctx.JSONP(http.StatusOK, gin.H{
		"book": books,
	})
}

func BookUpdate(ctx *gin.Context) {
	DB, _ := sql.Open("mysql", "test_u:test_pw@tcp(mysql:3306)/todo_db?charset=utf8")
	book := models.Book{}
	err := ctx.Bind(&book)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	_, err = book.Update(ctx, DB, boil.Infer())
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Server Error")
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookDelete(ctx *gin.Context) {
	DB, _ := sql.Open("mysql", "test_u:test_pw@tcp(mysql:3306)/todo_db?charset=utf8")
	book := models.Book{}
	err := ctx.Bind(&book)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	_, err = book.Delete(ctx, DB)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Server Error")
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func newBookRes(book *models.Book) *bookRes {
	return &bookRes{Book: book}
}

type bookRes struct {
	*models.Book
}
