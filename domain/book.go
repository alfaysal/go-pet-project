package domain

import "github.com/alfaysal/go-pet-project/dto"

type BookRepository interface {
	GetBookList(ctr *dto.BookCriteria) ([]Book, error)
	GetBook(id int) (Book, error)
	//CreateBook(w http.ResponseWriter, r *http.Request)
	//UpdateBook(w http.ResponseWriter, r *http.Request)
	//DeleteBook(w http.ResponseWriter, r *http.Request)
}

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type BookUsecase interface {
	GetBookList(ctr *dto.BookCriteria) ([]Book, error)
	GetBook(id int) (Book, error)
}
