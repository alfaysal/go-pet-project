package delivery

import (
	"github.com/alfaysal/go-pet-project/domain"
	"github.com/alfaysal/go-pet-project/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type BookHandler struct {
	bookUsecase domain.BookUsecase
}

func New(chi *chi.Mux, bookUseCase domain.BookUsecase) {
	bookHandler := &BookHandler{
		bookUsecase: bookUseCase,
	}

	chi.Get("/books", bookHandler.GetBookList)
}

func (bookHandler *BookHandler) GetBookList(w http.ResponseWriter, r *http.Request) {
	resp := utils.New(w)

	books, err := bookHandler.bookUsecase.GetBookList()

	if err != nil {
		resp.WithInternalServerError()
	}

	resp.WithData(&utils.Response{
		Data:    books,
		Message: "Fetch succesfully",
	})
}
