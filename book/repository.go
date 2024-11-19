package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(id int) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{db}
}