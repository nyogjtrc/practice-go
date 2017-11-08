package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID       uint
	Name     string
	Age      int
	CreateAt time.Time
}

func main() {

	db, err := GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	user := User{
		Name: "hi",
		Age:  9,
	}

	db.Create(&user)

	db.DropTable(&User{})
}

func GetConnection() (*gorm.DB, error) {
	return gorm.Open("mysql", "root:cypress@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
}
