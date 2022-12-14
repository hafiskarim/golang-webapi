package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:admin12345@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Databaes connection error")
	}

	db.AutoMigrate(&book.Book{})

	// CREATE DATA
	// book := book.Book{}
	// book.Title = "Hobbits"
	// book.Description = "Adventure Book"
	// book.Price = 200000
	// book.Rating = 4
	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("===================")
	// }

	var books []book.Book
	err = db.Debug().Where("rating =?", 5).Find(&books).Error
	if err != nil {
		fmt.Println("===================")
		fmt.Println("Error get book record")
		fmt.Println("===================")
	}

	for _, b := range books {
		fmt.Println("title: ", b.Title)
		fmt.Println("book object: ", b)
	}

	router := gin.Default()

	v1 := router.Group("/v1") // API Versioning v1

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/book/:id/:title", handler.BookHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8888")
}
