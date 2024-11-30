package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"go-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	service book.Service
}

func NewHandler(service book.Service) *handler {
	return &handler{service}
}

func (h *handler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func (h *handler) Getbook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid book ID",
		})
		return
	}
	
	book, err := h.service.FindBookByID(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *handler) GetBookbyQuery(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	price := c.Query("price")

	books, err := h.service.FindAllBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	filteredBooks := []book.Book{}
	for _, b := range books {
		if title != "" {
			if b.Title == title {
				filteredBooks = append(filteredBooks, b)
			}
		}
		if author != "" {
			if b.Author == author {
				filteredBooks = append(filteredBooks, b)
			}
		}
		if price != "" {
			if strconv.Itoa(b.Price) == price {
				filteredBooks = append(filteredBooks, b)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": filteredBooks,
	})
}



func (h *handler) PostBook(c *gin.Context) {
	var input book.Book
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

	book, err := h.service.CreateBook(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}



	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}


func (h *handler) GetBookAll(c *gin.Context) {
	books, err := h.service.FindAllBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}