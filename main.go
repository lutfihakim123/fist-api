package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	router.GET("/hello",helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBooksHandler)
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

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,numeric"`
}

func postBooksHandler(c *gin.Context){
	var bookInput BookInput


	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on field  %s, conditon: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title" : bookInput.Title,
		"price" : bookInput.Price,
	})
	
}