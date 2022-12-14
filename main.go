package main

import (
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

	bookRepository := book.NewRepository(db)
	book := book.Book{
		Title:       "Lord of The Rings Revision Edition",
		Description: "Dark Fantasy Book",
		Price:       250000,
		Rating:      4,
	}
	bookRepository.Create(book)

	// ===========
	// CREATE DATA
	// ===========
	// book := book.Book{}
	// book.Title = "Harry Potter"
	// book.Description = "Magic Book"
	// book.Price = 100000
	// book.Rating = 5
	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }

	// ========
	// GET DATA
	// ========
	// var book book.Book
	// err = db.Debug().Where("id =?", 1).Find(&book).Error
	// if err != nil {
	// 	fmt.Println("=====================")
	// 	fmt.Println("Error get book record")
	// 	fmt.Println("=====================")
	// }

	// for _, b := range books {
	// 	fmt.Println("title: ", b.Title)
	// 	fmt.Println("book object: ", b)
	// }

	// ===========
	// UPDATE DATA
	// ===========
	// book.Title = "Harry Potter (Revised edition)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("Error update book record")
	// 	fmt.Println("========================")
	// }

	// ===========
	// DELETE DATA
	// ===========
	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("Error delete book record")
	// 	fmt.Println("========================")
	// }

	router := gin.Default()

	v1 := router.Group("/v1") // API Versioning v1

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/book/:id/:title", handler.BookHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8888")
}
