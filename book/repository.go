package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(book Book) (Book, error)
	Update(id int, book Book) (Book, error)
	Delete(id int) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	if err != nil {
		return books, err
	}
	return books, nil
}

func (r repository) FindByID(id int) (Book, error) {
	var book Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r repository) Update(id int, book Book) (Book, error) {
	err := r.db.Save(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r repository) Delete(id int) (Book, error) {
	var book Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return book, err
	}
	err = r.db.Delete(&book, id).Error
	if err != nil {
		return book, err
	}
	return book, nil
}