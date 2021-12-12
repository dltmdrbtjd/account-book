package services

import "api/models"

type NewCategory struct {
	Title string `json:"title" binding:"requried"`
	Count string `json:"count" binding:"requried"`
}

func FindCategory() []models.Category {
	var category []models.Category
	models.DB.Find(&category)

	return category
}

func CreateCategory() NewCategory {
	var newCategory NewCategory

	return newCategory
}
