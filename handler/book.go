package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-api/book"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func Getbook(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "You requested to get a book with id: " + id,
	})
}

func GetBookbyQuery(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	c.JSON(http.StatusOK, gin.H{
		"message": "You requested to get a book with title: " + title + " and author: " + author,
	})
}



func PostBook(c *gin.Context) {
	var input book.BookInput
	err := c.ShouldBindJSON(&input)
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
		"title":  input.Title,
		"author": input.Author,
		"price":  input.Price,
	})
}
