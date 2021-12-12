package controllers

import (
	"api/models"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindCategories(c *gin.Context) {
	category := services.FindCategory()

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func CreateCategory(newCategory string) {
	var category models.Category
	if err := models.DB.Where("title = ?", newCategory).First(&category).Error; err != nil {
		newCategory := models.Category{Title: newCategory, Count: 1}
		models.DB.Create(&newCategory)
	}

	models.DB.Model(&category).Updates(map[string]interface{}{"count": category.Count + 1})
}

func UpdateCategory(c *gin.Context, currentCategory string, newCategory string) {
	if newCategory != "" {
		var category models.Category
		if err := models.DB.Where("title = ?", currentCategory).First(&category).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
			return
		}
		models.DB.Model(&category).Updates(map[string]interface{}{"count": category.Count - 1})
	}
}
