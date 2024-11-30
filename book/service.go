package book


type Service interface {
	FindAllBook() ([]Book, error)
	FindBookByID(id int) (Book, error)
	CreateBook(book Book) (Book, error)
	UpdateBookByID(id int, book Book) (Book, error)
	DeleteBookByID(id int) (Book, error)
}

type service struct {
	repository Repository
}	

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s service) FindAllBook() ([]Book, error) {
	books, err := s.repository.FindAll()
	if err != nil {
		return books, err
	}
	return books, nil
}

func (s service) FindBookByID(id int) (Book, error) {
	book, err := s.repository.FindByID(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (s service) CreateBook(book Book) (Book, error) {
	createdBook, err := s.repository.Create(book)
	if err != nil {
		return createdBook, err
	}
	return createdBook, nil
}

func (s service) UpdateBookByID(id int, book Book) (Book, error) {
	updatedBook, err := s.repository.Update(id, book)
	if err != nil {
		return updatedBook, err
	}
	return updatedBook, nil
}

func (s service) DeleteBookByID(id int) (Book, error) {
	book, err := s.repository.Delete(id)
	if err != nil {
		return book, err
	}
	return book, nil
}