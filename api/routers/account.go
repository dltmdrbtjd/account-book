package routers

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func AccountRoutes(route *gin.Engine) {
	route.GET("/account", controllers.FindAccounts)
}
