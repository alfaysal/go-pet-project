package delivery

import (
	"fmt"
	"github.com/alfaysal/go-pet-project/domain"
	"github.com/alfaysal/go-pet-project/dto"
	"github.com/alfaysal/go-pet-project/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type BookHandler struct {
	bookUsecase domain.BookUsecase
}

func New(chi *chi.Mux, bookUseCase domain.BookUsecase) {
	bookHandler := &BookHandler{
		bookUsecase: bookUseCase,
	}

	chi.Get("/books", bookHandler.GetBookList)
	chi.Get("/books/{id}", bookHandler.GetBook)
}

func (bookHandler *BookHandler) GetBookList(w http.ResponseWriter, r *http.Request) {
	resp := utils.New(w)

	bookDto := dto.BookCriteria{}

	if r.URL.Query().Get("id") != "" {
		bookID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)

		if err != nil {
			resp.WithError(&utils.Response{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}

		bookDto.ID = bookID
	}

	books, err := bookHandler.bookUsecase.GetBookList(&bookDto)

	if err != nil {
		resp.WithInternalServerError()
	}

	resp.WithData(&utils.Response{
		Data:    books,
		Message: "Fetch succesfully",
	})
}

func (bookHandler *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	resp := utils.New(w)
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		fmt.Println(err.Error())
	}

	book, err := bookHandler.bookUsecase.GetBook(id)

	resp.WithData(&utils.Response{
		Data:    book,
		Message: "Fetch succesfully",
	})
}
