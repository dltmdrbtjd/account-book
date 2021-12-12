package services

import "api/models"

type NewCategory struct {
	Title string `json:"title" binding:"requried"`
	Count int32  `json:"count" binding:"requried"`
}

type EditCategory struct {
	Title string `json:"title"`
	Count int32  `json:"count"`
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

func UpdateCategory() EditCategory {
	var updateCategory EditCategory

	return updateCategory
}
