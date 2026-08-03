package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"golang_clean_architecture/dto/product"
	"golang_clean_architecture/helper"
	"golang_clean_architecture/service"
	"golang_clean_architecture/web"
)

type ProductHandler struct {
	Service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (handler *ProductHandler) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	products := handler.Service.FindAll(ctx)

	responsse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   products,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(responsse)
	helper.PanicIfError(err)
}

func (handler *ProductHandler) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()
	id, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	product, err := handler.Service.FindById(ctx, id)
	if err != nil {
		responsse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(responsse)
		helper.PanicIfError(err)
		return
	}

	responsse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responsse)
	helper.PanicIfError(err)
}

func (handler *ProductHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)
	productCreateRequest := product.ProductCreateRequest{}
	err := decoder.Decode(&productCreateRequest)
	helper.PanicIfError(err)

	product, err := handler.Service.Create(ctx, productCreateRequest)
	helper.PanicIfError(err)

	responsse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   product,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responsse)
	helper.PanicIfError(err)
}
