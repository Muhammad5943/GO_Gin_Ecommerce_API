package dtos

import "github.com/muhammad5943/go_gin_ecommerce_api/models"

func CreateHomeResponse(tags []models.Tag, categories []models.Category) map[string]interface{} {
	return CreateSuccessDto(map[string]interface{}{
		"tags":       CreateTagListDto(tags),
		"categories": CreateCategoryListDto(categories),
	})
}
