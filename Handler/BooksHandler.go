package Handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main.go/Book"
)


func RootHandler (c *gin.Context) {
	c.JSON(http.StatusOK ,gin.H {
		"name" : "Maula Izza Aziiz",
		"Bio" : "Software Engineer",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"title" : "Hello World",
		"subtitle" : "Belajar Golang bersama ku",
	})
}

// path id
func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Query("Title")

	c.JSON(http.StatusOK, gin.H {
		"id" : id,
		"nama" : title,
	})
}

// 127.0.0.1:8080/query?price=40&title=Cantik itu luka
func UseQuery(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H {
		"title" : title,
		"price" : price,
	})
}

func UseDoubleQuery(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H {
		"id" : id,
		"title" : title,
	})
}

func PostBookHandler(c *gin.Context) {
	var bookInput Book.InputBook

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf ("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title" : bookInput.Title,
		"price" : bookInput.Price,
	})
}

func GetPostBookHandler (c *gin.Context) {
	var bookInput Book.InputBook
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusOK, gin.H{"errors" : errorMessages,})
		return
	}
}