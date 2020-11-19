package model

type Book struct {
    Id    int  `xorm:"pk autoincr int" form:"id" json:"id"`
    Title string `xorm:"varchar(255)" json:"title" form:"title"`
    Content string `xorm:"varchar(255)" json:"content" form:"content"`
}