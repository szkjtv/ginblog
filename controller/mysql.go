package controller

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Dbinit() (db *gorm.DB) { //tests   world  golang
	
	db, err = gorm.Open("mysql", "数据库名:密码@tcp(数据库ip:3306)/用户名?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Aritclee{})
	db.AutoMigrate(&LoginUser{})

	return db

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Aritclee struct {
	gorm.Model
	Title   string `gorm:"type:longtext"`
	Content string `gorm:"type:longtext"`
	UserID  int    //用来存取登录用户的id
	// User    []LoginUser
}

type LoginUser struct {
	gorm.Model
	Username string
	Account  string
	Password string
	// Ariclt   []Aritclee
}
