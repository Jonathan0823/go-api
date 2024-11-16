package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "fmt"
)

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	  "message": "Hello World",
	})
  })
  fmt.Println("Server is running on port 8080")
  http.ListenAndServe("localhost:8080", r)
}