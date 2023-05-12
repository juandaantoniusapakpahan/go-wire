package middleware

import (
	"net/http"

	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
	"github.com/juandaantoniusapakpahan/go-restful-api/model/web"
)

type Middelware struct {
	Handler http.Handler
}

func NewMiddleare(handler http.Handler) *Middelware {
	return &Middelware{Handler: handler}
}

func (middleware *Middelware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "SECRETKEY" == r.Header.Get("Z-API-KEY") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.Response{
			Code:   http.StatusUnauthorized,
			Status: "UnAuthorized",
		}
		helper.WriteToResponseBody(w, webResponse)
	}
}
