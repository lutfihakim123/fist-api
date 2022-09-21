package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	router.GET("/hello",helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.Run()
}


func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"Name" : "Lutfi Rahman Hakim",
	})
}

func helloHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"-_-" : "Hello Lutfi",
	})
}

func booksHandler(c *gin.Context)  {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id" : id,
		"title" : title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title" : title,
		"price" : price,
	})
}

