package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	// database, err := gorm.Open("sqlite3", "test.db")

	// if err != nil {
	// 	panic("Failed to connect to database!")
	// }

	user := "postgres"
	password := "123456"
	database := "golangdb"
	// host := "localhost"
	port := "5432"
	// user=gorm password=gorm dbname=gorm port=9920 sslmode=disable
	prosgret_conname := fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=disable", user, password, database, port)
	fmt.Println("conname is\t\t", prosgret_conname)
	db, err := gorm.Open("postgres", prosgret_conname)
	// println(err.Error())
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{})

	DB = db
}
