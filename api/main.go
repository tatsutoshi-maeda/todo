package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/go-sql-driver/mysql"
	"api/controller"
	"api/middleware"
)

func main() {
	engine := gin.Default()
	// ミドルウェア
    config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	engine.Use(cors.New(config))
	engine.Use(middleware.RecordUaAndTime)
	bookEngine := engine.Group("/book")
	{
		v1 := bookEngine.Group("/v1")
		{
			v1.POST("/add", controller.BookAdd)
			v1.GET("/list", controller.BookList)
			v1.PUT("/update", controller.BookUpdate)
			v1.DELETE("/delete", controller.BookDelete)
		}
	}
	engine.Run(":8080")
}
