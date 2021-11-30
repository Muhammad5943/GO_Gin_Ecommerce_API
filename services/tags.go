package services

import (
	"github.com/muhammad5943/go_gin_ecommerce_api/infrastructure"
	"github.com/muhammad5943/go_gin_ecommerce_api/models"
)

func FetchAllTags() ([]models.Tag, error) {
	database := infrastructure.GetDb()
	var tags []models.Tag
	err := database.Preload("Images", "tag_id IS NOT NULL").Find(&tags).Error
	return tags, err
}
