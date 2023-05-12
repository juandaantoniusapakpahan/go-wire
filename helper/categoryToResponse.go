package helper

import (
	"github.com/juandaantoniusapakpahan/go-restful-api/model/domain"
	"github.com/juandaantoniusapakpahan/go-restful-api/model/web"
)

func CategoryToResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func CategoriesToResponses(categories []domain.Category) []web.CategoryResponse {
	var responsesCategory []web.CategoryResponse
	for _, category := range categories {
		responsesCategory = append(responsesCategory, CategoryToResponse(category))
	}
	return responsesCategory
}
