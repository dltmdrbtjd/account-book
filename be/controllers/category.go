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
	category := services.CreateCategory()
	if err := models.DB.Where("title = ?", newCategory).First(&category).Error; err != nil {
		newCategory := models.Category{Title: newCategory, Count: 1}
		models.DB.Create(&newCategory)
		return
	}

	models.DB.Model(&category).Updates(map[string]interface{}{"count": category.Count + 1})
}

func UpdateCategory(c *gin.Context, currentCategory string, newCategory string) {
	if newCategory != "" {
		category := services.UpdateCategory()
		if err := models.DB.Where("title = ?", currentCategory).First(&category).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
			return
		}
		models.DB.Model(&category).Updates(map[string]interface{}{"count": category.Count - 1})
	}
}

func DeleteCategory(c *gin.Context, title string) {
	var category models.Category
	if err := models.DB.Where("title = ?", title).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	if category.Count == 1 {
		models.DB.Where("title = ?", title).Delete(&category)
		return
	}
	models.DB.Model(&category).Updates(map[string]interface{}{"count": category.Count - 1})
}
