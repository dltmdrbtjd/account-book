package controllers

import (
	"api/models"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAccounts(c *gin.Context) {
	accounts := services.FindAccounts()

	c.JSON(http.StatusOK, gin.H{"data": accounts})
}

func FindAccount(c *gin.Context) {
	var account models.Account
	if err := models.DB.Where("id = ?", c.Param("id")).First(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": account})
}

func CreateAccount(c *gin.Context) {
	input := services.CreateAccount()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CreateCategory(input.Category)

	newAccount := models.Account{Title: input.Title, Category: input.Category, Money: input.Money, CreateAt: input.CreateAt}
	models.DB.Create(&newAccount)

	c.JSON(http.StatusOK, gin.H{"data": newAccount})
}

func UpdateAccount(c *gin.Context) {
	var account models.Account
	if err := models.DB.Where("id = ?", c.Param("id")).First(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	input := services.UpdateAccount()
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Title == "" {
		input.Title = account.Title
	}

	UpdateCategory(c, account.Category, input.Category)

	if input.Category == "" {
		input.Category = account.Category
	}
	if input.Money < 1 {
		input.Money = account.Money
	}

	CreateCategory(input.Category)

	models.DB.Model(&account).Updates(map[string]interface{}{"title": input.Title, "Category": input.Category, "Money": input.Money, "UpdateAt": input.UpdateAt})

	c.JSON(http.StatusOK, gin.H{"data": account})
}

func DeleteAccount(c *gin.Context) {
	var account models.Account
	if err := models.DB.Where("id = ?", c.Param("id")).First(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&account)
	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
