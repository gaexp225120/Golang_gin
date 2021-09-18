package middleware

import (
	// "github.com/go-pg/pg/v9"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id   string `json:"User_Id" gorm:"primary_key"`
	Name string `json:"User_Name" `
}

var db *gorm.DB

// 連接DB
func Connect() *gorm.DB {
	dsn := "sslmode=disable host=127.0.0.1 port=5432 user=user1 password=pwd dbname=test"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
	db = database
	db.AutoMigrate(&User{})
	return db
}
