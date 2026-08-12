package route

import (
	controller "golang_clean_architecture/internal/delivery/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(bookController controller.BookController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/books", bookController.FindAll)
	router.GET("/api/books/:bookId", bookController.FindById)
	router.POST("/api/books", bookController.Create)
	router.PUT("/api/books/:bookId", bookController.Update)
	router.DELETE("/api/books/:bookId", bookController.Delete)

	return router
}
