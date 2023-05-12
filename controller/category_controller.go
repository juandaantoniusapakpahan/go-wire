package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(w http.ResponseWriter, r *http.Request, paramas httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, paramas httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, paramas httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, paramas httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, paramas httprouter.Params)
}
