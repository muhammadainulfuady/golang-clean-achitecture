package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang_clean_architecture/internal/model"
	"golang_clean_architecture/internal/usecase"

	"github.com/julienschmidt/httprouter"
)

type BookControllerImpl struct {
	BookUsecase usecase.BookUsecase
}

func NewBookController(bookUsecase usecase.BookUsecase) BookController {
	return &BookControllerImpl{
		BookUsecase: bookUsecase,
	}
}

func (controller *BookControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createBookRequest := model.CreateBookRequest{}
	err := json.NewDecoder(request.Body).Decode(&createBookRequest)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD REQUEST",
			Message: err.Error(),
		})
		return
	}

	bookResponse, err := controller.BookUsecase.Create(request.Context(), createBookRequest)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL SERVER ERROR",
			Message: err.Error(),
		})
		return
	}

	webResponse := model.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success Create Book",
		Data:    bookResponse,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(webResponse)
}

func (controller *BookControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateBookRequest := model.UpdateBookRequest{}
	err := json.NewDecoder(request.Body).Decode(&updateBookRequest)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD REQUEST",
			Message: err.Error(),
		})
		return
	}

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD REQUEST",
			Message: "Invalid Book ID",
		})
		return
	}

	updateBookRequest.ID = id

	bookResponse, err := controller.BookUsecase.Update(request.Context(), updateBookRequest)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL SERVER ERROR",
			Message: err.Error(),
		})
		return
	}

	webResponse := model.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success Update Book",
		Data:    bookResponse,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(webResponse)
}

func (controller *BookControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD REQUEST",
			Message: "Invalid Book ID",
		})
		return
	}

	err = controller.BookUsecase.Delete(request.Context(), id)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL SERVER ERROR",
			Message: err.Error(),
		})
		return
	}

	webResponse := model.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success Delete Book",
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(webResponse)
}

func (controller *BookControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD REQUEST",
			Message: "Invalid Book ID",
		})
		return
	}

	bookResponse, err := controller.BookUsecase.FindById(request.Context(), id)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(model.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: err.Error(),
		})
		return
	}

	webResponse := model.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success Get Book By ID",
		Data:    bookResponse,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(webResponse)
}

func (controller *BookControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookResponses, err := controller.BookUsecase.FindAll(request.Context())
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL SERVER ERROR",
			Message: err.Error(),
		})
		return
	}

	webResponse := model.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success Get All Books",
		Data:    bookResponses,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(webResponse)
}
