package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"

	"github.com/juandaantoniusapakpahan/go-restful-api/app"
	"github.com/juandaantoniusapakpahan/go-restful-api/controller"
	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
	"github.com/juandaantoniusapakpahan/go-restful-api/middleware"
	"github.com/juandaantoniusapakpahan/go-restful-api/model/domain"
	"github.com/juandaantoniusapakpahan/go-restful-api/repository"
	"github.com/juandaantoniusapakpahan/go-restful-api/services"
	"github.com/stretchr/testify/assert"
)

func setRouter(DB *sql.DB) http.Handler {
	db := NewConnect()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewMiddleare(router)
}

func NewConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:@root123@tcp(localhost:3306)/belajar_golang_result_api_test")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}

func DeleteCategory(DB *sql.DB) {
	_, err := DB.ExecContext(context.Background(), "DELETE FROM category where 1=1")
	helper.PanicIfError(err)
}

type M map[string]interface{}

func TestCreateCategorySuccess(t *testing.T) {
	DB := NewConnect()
	router := setRouter(DB)
	DeleteCategory(DB)

	dataJson := strings.NewReader(`{"name":"Sumsang"}`)
	request := httptest.NewRequest("POST", "http://localhost:8080/api/categories", dataJson)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Z-API-KEY", "SECRETKEY")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	responseBody := M{}
	body, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(body, &responseBody)

	// Assert
	assert.Equal(t, 200, recorder.Result().StatusCode)
	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Sumsang", responseBody["data"].(map[string]interface{})["name"])

}

func TestCreateCategoryFailed(t *testing.T) {
	DB := NewConnect()
	router := setRouter(DB)
	DeleteCategory(DB)

	dataJson := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest("POST", "http://localhost:8080/api/categories", dataJson)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Z-API-KEY", "SECRETKEY")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	responseBody := M{}
	body, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, recorder.Result().StatusCode)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	DB := NewConnect()
	router := setRouter(DB)
	DeleteCategory(DB)

	categoryPayload := domain.Category{
		Name: "GGWP",
	}

	tx, _ := DB.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, categoryPayload)
	tx.Commit()

	dataJson := strings.NewReader(`{"name":"LUXEX"}`)
	request := httptest.NewRequest("PUT", "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), dataJson)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Z-API-KEY", "SECRETKEY")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	responseBody := M{}
	body, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, recorder.Result().StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))

}

func TestUpdateCategoryFailed(t *testing.T) {
	DB := NewConnect()
	route := setRouter(DB)
	DeleteCategory(DB)

	categoryPayload := domain.Category{
		Name: "Had updated",
	}

	tx, _ := DB.Begin()
	respositoryCategory := repository.NewCategoryRepository()
	category := respositoryCategory.Save(context.Background(), tx, categoryPayload)
	tx.Commit()

	dataJson := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest("PUT", "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), dataJson)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Z-API-KEY", "SECRETKEY")
	recorder := httptest.NewRecorder()

	route.ServeHTTP(recorder, request)

	responseBody := M{}
	body, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, recorder.Result().StatusCode)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	DB := NewConnect()
	router := setRouter(DB)
	DeleteCategory(DB)

	categoryPayload := domain.Category{
		Name: "GGWP",
	}

	tx, _ := DB.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, categoryPayload)
	tx.Commit()

	dataJson := strings.NewReader(`{"name":"LUXEX"}`)
	request := httptest.NewRequest("DELETE", "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), dataJson)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Z-API-KEY", "SECRETKEY")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	responseBody := M{}
	body, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, recorder.Result().StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	//assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))

}

func TestDeleteCategoryWithNotFoundID(t *testing.T) {
	DB := NewConnect()
	router := setRouter(DB)
	DeleteCategory(DB)

	categoryPayload := domain.Category{
		Name: "GGWP",
	}

	tx, _ := DB.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, categoryPayload)
	tx.Commit()

	dataJson := strings.NewReader(`{"name":"LUXEX"}`)
	request := httptest.NewRequest("DELETE", "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id+1), dataJson)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Z-API-KEY", "SECRETKEY")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	responseBody := M{}
	body, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, recorder.Result().StatusCode)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])

}

func TestGetCategorySuccess(t *testing.T) {
	DB := NewConnect()
	router := setRouter(DB)
	DeleteCategory(DB)

	categoryPayload := domain.Category{
		Name: "GGWP",
	}

	tx, _ := DB.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, categoryPayload)
	tx.Commit()

	dataJson := strings.NewReader(`{"name":"LUXEX"}`)
	request := httptest.NewRequest("GET", "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), dataJson)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Z-API-KEY", "SECRETKEY")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	responseBody := M{}
	body, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, recorder.Result().StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestGetCategoryNotFound(t *testing.T) {
	DB := NewConnect()
	router := setRouter(DB)
	DeleteCategory(DB)

	categoryPayload := domain.Category{
		Name: "GGWP",
	}

	tx, _ := DB.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, categoryPayload)
	tx.Commit()

	dataJson := strings.NewReader(`{"name":"LUXEX"}`)
	request := httptest.NewRequest("GET", "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id+1), dataJson)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Z-API-KEY", "SECRETKEY")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	responseBody := M{}
	body, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, recorder.Result().StatusCode)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestGetAllCategorySuccess(t *testing.T) {
	DB := NewConnect()
	router := setRouter(DB)
	DeleteCategory(DB)

	tx, _ := DB.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category1 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Test1",
	})
	category2 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Test2",
	})

	tx.Commit()

	dataJson := strings.NewReader(`{"name":"LUXEX"}`)
	request := httptest.NewRequest("GET", "http://localhost:8080/api/categories", dataJson)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Z-API-KEY", "SECRETKEY")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	responseBody := M{}
	body, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(body, &responseBody)

	categories1 := responseBody["data"].([]interface{})[0].(map[string]interface{})
	categories2 := responseBody["data"].([]interface{})[1].(map[string]interface{})

	assert.Equal(t, 200, recorder.Result().StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, category1.Id, int(categories1["id"].(float64)))
	assert.Equal(t, category2.Id, int(categories2["id"].(float64)))

}
