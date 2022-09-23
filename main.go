package main

import (
	"fmt"
	"log"
	"pustaka-api/book"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db Connection Error")
	}

	bookRepository := book.NewRepository(db)


	// Find All Data

	books, err := bookRepository.FindAll()

	for _, book := range books {
		fmt.Println("Title: ", book.Title)
	} 


	// Find By id

	// book, err := bookRepository.FindById(1)

	// fmt.Println("Title: ", book.Title)

	// create data 

	// book := book.Book{
	// 	Title:  "Buku Kedua",
	// 	Description: "Buku Kedua Saya",
	// 	Price: 60000,
	// 	Rating: 4,
	// 	Discount: 0,
	// }

	// bookRepository.Create(book)

	// router := gin.Default()
	// v1 := router.Group("/v1")
	// v1.GET("/", handler.RootHandler)
	// v1.GET("/hello",handler.HelloHandler)
	// v1.GET("/books/:id/:title", handler.BooksHandler)
	// v1.GET("/query", handler.QueryHandler)
	// v1.POST("/books", handler.PostBooksHandler)
	// router.Run()
}