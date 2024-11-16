package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "fmt"
)

func main() {
  r := gin.Default()
  r.GET("/", rootHandler)

  fmt.Println("Server is running at localhost:8080")
  http.ListenAndServe("localhost:8080", r)
}

func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
	"message": "Hello World",
  })
}

