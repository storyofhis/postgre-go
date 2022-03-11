package main

import (
	"github.com/gin-gonic/gin"
	"main.go/Handler"
)

const (
    dbDriver      = "postgres"
    dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
    serverAddress = "0.0.0.0:8080"
)


func main() {
	router := gin.Default()

	// API Versioning
	v1 := router.Group("/v1")

	// HTTP GET METHOD
	v1.GET("/", Handler.RootHandler)
	router.GET("/hello", Handler.HelloHandler)
	router.GET("/books/:id", Handler.BooksHandler)
	router.GET("/books", Handler.UseQuery)
	router.GET("/books/:id/:title", Handler.UseDoubleQuery)

	// HTTP 

	router.Run()
}

