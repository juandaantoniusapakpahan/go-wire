package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/juandaantoniusapakpahan/go-restful-api/app"
	"github.com/juandaantoniusapakpahan/go-restful-api/controller"
	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
	"github.com/juandaantoniusapakpahan/go-restful-api/middleware"
	"github.com/juandaantoniusapakpahan/go-restful-api/repository"
	"github.com/juandaantoniusapakpahan/go-restful-api/services"
)

func main() {
	db := app.Connect()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	middleware := middleware.NewMiddleare(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware,
	}

	fmt.Println("Server started at: localhost:8080")
	err := server.ListenAndServe()
	helper.ErrorHandle(err)
}
