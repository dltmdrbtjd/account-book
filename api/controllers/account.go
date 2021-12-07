package controllers

import (
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAccounts(c *gin.Context) {
	accounts := services.FindAccounts()

	c.JSON(http.StatusOK, gin.H{"data": accounts})
}
