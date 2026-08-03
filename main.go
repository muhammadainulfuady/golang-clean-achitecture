package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"golang_clean_architecture/database"
	"golang_clean_architecture/exception"
	"golang_clean_architecture/handler"
	"golang_clean_architecture/helper"
	"golang_clean_architecture/repository"
	"golang_clean_architecture/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.NewDB()

	productRepository := repository.ProductRepositoryImpl{}
	productService := service.NewProductService(&productRepository, db)

	productHandler := handler.NewProductHandler(productService)

	router := httprouter.New()
	router.GET("/api/products/", productHandler.FindAll)
	router.GET("/api/products/:id", productHandler.FindById)
	router.POST("/api/products/", productHandler.Create)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

	log.Println("Server running on port http://localhost:3000")
}
