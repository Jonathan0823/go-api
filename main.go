package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-api/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  "go-api/book"
)

func main() {
  loadEnv := godotenv.Load()
  if loadEnv != nil {
    log.Fatal("Error loading .env file")
  }

  
  dburl := os.Getenv("DATABASE_URL")
  if dburl == "" {
    log.Fatal("DATABASE_URL is not set")
	}
  
  db, err := gorm.Open(postgres.Open(dburl), &gorm.Config{})
  if err != nil {
    log.Fatal("failed to connect database")
  }
  
  log.Println("Database connection established")
  
  db.AutoMigrate(&book.Book{})


  bookRepository := book.NewRepository(db)
  bookService := book.NewService(bookRepository)
  bookHandler := handler.NewHandler(bookService)

  

  r := gin.Default()
  
  v1 := r.Group("/v1")
    v1.GET("/", )
    v1.GET("/books/:id", bookHandler.Getbook)
    v1.GET("/books", bookHandler.GetBookbyQuery)
    v1.POST("/books", bookHandler.PostBook)
    v1.GET("/books/all", bookHandler.GetBookAll)

  fmt.Println("Server is running at localhost:8080")
  http.ListenAndServe("localhost:8080", r)
}
