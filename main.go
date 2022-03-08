package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Person struct {
	gorm.Model

	Name string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model

	Title string
	Author string
	CallNumber int `gorm:"unique_index"`
	PersonID int
}

var db *gorm.DB
var err error

func main() {
	// load environtment variables
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// Database Connetion Strings
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	// Openning Connection to Database
	db, err := gorm.Open(dialect, dbURI);
	
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("Successfully connected to database")
	}

	// close main function
	// defer db.Close()

	// database miggration 
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})
}