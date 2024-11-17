package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
  r := gin.Default()
  r.GET("/", rootHandler)
  r.GET("/books/:id", getbook)
  r.GET("/books", getBookbyQuery)
  r.POST("/books", postBook)

  fmt.Println("Server is running at localhost:8080")
  http.ListenAndServe("localhost:8080", r)
}

func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
	"message": "Hello World",
  })
}

func getbook(c *gin.Context){
  id := c.Param("id")
  c.JSON(http.StatusOK, gin.H{
    "message": "You requested to get a book with id: " + id,
  })
}

func getBookbyQuery(c *gin.Context){
  title := c.Query("title")
  author := c.Query("author")
  c.JSON(http.StatusOK, gin.H{
    "message": "You requested to get a book with title: " + title + " and author: " + author,
  })
}

type BookInput struct {
  Title string `json:"title" binding:"required"`
  Author string `json:"author" binding:"required"`
  Price json.Number `json:"price" binding:"required,number"`
}

func postBook(c *gin.Context){
 var input BookInput
 err:= c.ShouldBindJSON(&input)
  if err != nil {
  
    var ve validator.ValidationErrors
    if errors.As(err, &ve) {
    errorMessages := []string{}
      for _, e := range err.(validator.ValidationErrors) {
        errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
        errorMessages = append(errorMessages, errorMessage)
    }
      c.JSON(http.StatusBadRequest, gin.H{
        "error": errorMessages,
      })
      return
    }
    
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
}
  

  c.JSON(http.StatusOK, gin.H{
    "title": input.Title,
    "author": input.Author,
    "price": input.Price,
  })
}

