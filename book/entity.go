package book

import "time"

type Book struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    int    `json:"price"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}