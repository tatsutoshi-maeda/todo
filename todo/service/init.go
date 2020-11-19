package service

import (
	"errors"
	"fmt"
	"log"
	"todo/model"

	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := "test_u:test_pw@(127.0.0.1:3307)/todo_db?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.Book))
	fmt.Println("init data base ok")
}
