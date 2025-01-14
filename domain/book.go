package domain

type BookRepository interface {
	GetBookList() ([]Book, error)
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
	GetBookList() ([]Book, error)
}
