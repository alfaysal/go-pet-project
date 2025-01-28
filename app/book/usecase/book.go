package usecase

import (
	"github.com/alfaysal/go-pet-project/domain"
	"github.com/alfaysal/go-pet-project/dto"
	"github.com/alfaysal/go-pet-project/internal/cache"
)

type BookUsecase struct {
	bookRepo domain.BookRepository
	cache    cache.Cache
}

func New(bookRepo domain.BookRepository, cache cache.Cache) domain.BookUsecase {
	return &BookUsecase{
		bookRepo: bookRepo,
		cache:    cache,
	}
}

func (bookUsecase *BookUsecase) GetBookList(ctr *dto.BookCriteria) ([]domain.Book, error) {
	books, err := bookUsecase.bookRepo.GetBookList(ctr)

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
