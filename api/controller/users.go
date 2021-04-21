package controller

import (
	"api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func UserAdd(ctx *gin.Context) {
	DB, _ := sql.Open("mysql", "test_u:test_pw@tcp(mysql:3306)/todo_db?charset=utf8")
	user := models.User{}
	err := ctx.Bind(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	err = user.Insert(ctx, DB, boil.Infer())
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Server Error")
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func UserList(ctx *gin.Context) {
	DB, _ := sql.Open("mysql", "test_u:test_pw@tcp(mysql:3306)/todo_db?charset=utf8")
	var err error
	users := []*models.User{}

	users, err = models.Users().All(ctx, DB)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Server Error")
		return
	}
	ctx.JSONP(http.StatusOK, gin.H{
		"user": users,
	})
}

func UserUpdate(ctx *gin.Context) {
	DB, _ := sql.Open("mysql", "test_u:test_pw@tcp(mysql:3306)/todo_db?charset=utf8")
	user := models.User{}
	err := ctx.Bind(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	_, err = user.Update(ctx, DB, boil.Infer())
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Server Error")
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func UserDelete(ctx *gin.Context) {
	DB, _ := sql.Open("mysql", "test_u:test_pw@tcp(mysql:3306)/todo_db?charset=utf8")
	user := models.User{}
	err := ctx.Bind(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	_, err = user.Delete(ctx, DB)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Server Error")
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func newUserRes(user *models.User) *userRes {
	return &userRes{User: user}
}

type userRes struct {
	*models.User
}
