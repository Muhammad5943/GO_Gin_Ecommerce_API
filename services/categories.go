package services

import (
	"github.com/muhammad5943/go_gin_ecommerce_api/infrastructure"
	"github.com/muhammad5943/go_gin_ecommerce_api/models"
)

func FetchAllCategories() ([]models.Category, error) {
	database := infrastructure.GetDb()
	var categories []models.Category
	err := database.Preload("Images", "category_id IS NOT NULL").Find(&categories).Error
	return categories, err
}
