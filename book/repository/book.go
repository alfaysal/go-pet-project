package repository

import (
	"fmt"
	"github.com/alfaysal/go-pet-project/domain"
	"gorm.io/gorm"
)

type BookSQL struct {
	db *gorm.DB
}

func (bookSql *BookSQL) GetBookList() ([]domain.Book, error) {
	var books []domain.Book

	if err := bookSql.db.Find(&books).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch books: %w", err)
	}
	fmt.Println(books)
	return books, nil
}

func New(db *gorm.DB) domain.BookRepository {
	return &BookSQL{
		db: db,
	}
}
