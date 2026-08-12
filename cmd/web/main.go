package main

import (
	"fmt"
	"net/http"

	"golang_clean_architecture/internal/config"
	delivery "golang_clean_architecture/internal/delivery/http"
	"golang_clean_architecture/internal/delivery/http/route"
	"golang_clean_architecture/internal/repository"
	"golang_clean_architecture/internal/usecase"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := config.NewDB()
	validate := validator.New()

	bookRepository := repository.NewBookRepository()
	bookUsecase := usecase.NewBookUsecase(db, bookRepository, validate)
	bookController := delivery.NewBookController(bookUsecase)

	router := route.NewRouter(bookController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Server running on http://localhost:3000")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
