package routers

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func Category(route *gin.Engine) {
	route.GET("/category", controllers.FindCategories)
}
