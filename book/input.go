package book

import "encoding/json"

type BookInput struct {
	Title  string      `json:"title" binding:"required"`
	Author string      `json:"author" binding:"required"`
	Price  json.Number `json:"price" binding:"required,number"`
}