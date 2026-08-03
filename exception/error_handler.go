package exception

import (
	"encoding/json"
	"net/http"

	"golang_clean_architecture/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")

	switch err.(type) {
	case *NotFoundError:
		response := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: err.(*NotFoundError).Error(),
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
	case *ValidationError:
		response := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.(*ValidationError).Messages,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	default:
		response := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
	}
}
