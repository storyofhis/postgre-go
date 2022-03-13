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
	v1.GET("/hello", Handler.HelloHandler)
	v1.GET("/books/:id", Handler.BooksHandler)
	v1.GET("/books/:id/:title", Handler.UseDoubleQuery)
	v1.GET("/query", Handler.UseQuery)

	// HTTP POST METHOD
	v1.POST("/books", Handler.PostBookHandler)
	
	router.Run()
}

