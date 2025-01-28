package repository

import (
	"fmt"
	"github.com/alfaysal/go-pet-project/domain"
	"github.com/alfaysal/go-pet-project/dto"
	"gorm.io/gorm"
)

type BookSQL struct {
	db *gorm.DB
}

func (bookSql *BookSQL) GetBookList(ctr *dto.BookCriteria) ([]domain.Book, error) {
	var books []domain.Book
	query := bookSql.db

	if ctr.ID != 0 {
		query = query.Where("id = ?", ctr.ID)
	}

	if err := query.Find(&books).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch books: %w", err)
	}

	return books, nil
}

func New(db *gorm.DB) domain.BookRepository {
	return &BookSQL{
		db: db,
	}
}

func (bookSql *BookSQL) GetBook(id int) (domain.Book, error) {
	var book domain.Book

	if err := bookSql.db.Find(&book, id).Error; err != nil {
		return domain.Book{}, fmt.Errorf("failed to fetch books: %w", err)
	}

	return book, nil
}
