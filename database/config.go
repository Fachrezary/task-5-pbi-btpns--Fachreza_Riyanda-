package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

// DB is the global variable for the database connection.
var DB *gorm.DB

// ConnectDB connects to the database.
func ConnectDB() {
	dbHost := os.Getenv("localhost")
	dbPort := os.Getenv("3306")
	dbUser := os.Getenv("root")
	dbPassword := os.Getenv("")
	dbName := os.Getenv("project_api")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	conn, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	DB = conn
}

// CloseDB closes the database connection.
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}