package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Product _
type Product struct {
	ID        uint   `gorm:"primary_key"`
	Code      string `gorm:"not null;unique_index"`
	Price     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	// Connect database
	db, err := gorm.Open("mysql", "root:root@tcp/practice?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.Debug().AutoMigrate(&Product{})
	defer db.DropTable(Product{})

	// Create
	err = db.Create(&Product{Code: "L1212", Price: 1000}).Error
	if err != nil {
		panic(err)
	}

	err = db.Create(&Product{Code: "L1234", Price: 1234}).Error
	if err != nil {
		panic(err)
	}

	// Read
	var products []Product
	var product Product

	// find product with id 1
	if err := db.Find(&products).Error; err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(products)

	// find product with code l1212
	if err := db.Where("code = ?", "L1212").First(&product).Error; err != nil {
		panic(err)
	}
	fmt.Println("product L1212", product)

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)
	fmt.Println("product", product)

	// Delete - delete product
	err = db.Debug().Delete(Product{}).Error
	if err != nil {
		panic(err)
	}

}
