package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/alfaysal/go-pet-project/domain"
	"github.com/alfaysal/go-pet-project/dto"
	"github.com/alfaysal/go-pet-project/internal/cache"
	"time"
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
	var books []domain.Book
	val, err := bookUsecase.cache.Get("books")

	if val != "" {
		err = json.Unmarshal([]byte(val), &books)

		if err != nil {
			return nil, err
		}
		fmt.Println("Cache hit")
		return books, nil
	}

	books, err = bookUsecase.bookRepo.GetBookList(ctr)

	if err != nil {
		return nil, err
	}

	if err := bookUsecase.cache.Set("books", books, time.Second*60); err != nil {
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
