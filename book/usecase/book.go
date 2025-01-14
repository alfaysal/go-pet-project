package usecase

import "github.com/alfaysal/go-pet-project/domain"

type BookUsecase struct {
	bookRepo domain.BookRepository
}

func New(bookRepo domain.BookRepository) domain.BookUsecase {
	return &BookUsecase{
		bookRepo: bookRepo,
	}
}

func (bookUsecase *BookUsecase) GetBookList() ([]domain.Book, error) {
	books, err := bookUsecase.bookRepo.GetBookList()

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (bookUsecase *BookUsecase) GetBook(id int) (domain.Book, error) {
	book, err := bookUsecase.bookRepo.GetBook(id)

	if err != nil {
		return domain.Book{}, err
	}

	return book, nil
}
