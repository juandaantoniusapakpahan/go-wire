package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
	"github.com/juandaantoniusapakpahan/go-restful-api/model/web"
)

func ErrorHalder(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}
	if badRequest(w, r, err) {
		return
	}

	internalError(w, r, err)
}
func badRequest(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exp, ok := err.(validator.ValidationErrors)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exp.Error(),
		}
		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exp, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.Response{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exp.Error,
		}
		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.Response{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	helper.WriteToResponseBody(w, webResponse)
}
