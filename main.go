package main

import (
	"fmt"
	"net/http"

	"go-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  v1 := r.Group("/v1")
    v1.GET("/", )
    v1.GET("/books/:id", handler.Getbook)
    v1.GET("/books", handler.GetBookbyQuery)
    v1.POST("/books", handler.PostBook)

  fmt.Println("Server is running at localhost:8080")
  http.ListenAndServe("localhost:8080", r)
}
