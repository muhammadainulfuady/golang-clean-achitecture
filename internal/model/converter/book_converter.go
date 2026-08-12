package converter

import (
	"golang_clean_architecture/internal/entity"
	"golang_clean_architecture/internal/model"
)

func ToBookConverter(book *entity.Book) model.BookResponse {
	return model.BookResponse{
		ID:              book.ID,
		Name:            book.Name,
		Author:          book.Author,
		PublicationYear: book.PublicationYear,
		Description:     book.Description,
	}
}
