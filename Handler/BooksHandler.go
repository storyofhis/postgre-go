package Handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, gin.H {
		"id" : id,
	})
}

func UseQuery(c *gin.Context) {
	title := c.Query("Title")
	price := c.Query("Price")

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