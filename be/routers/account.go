package routers

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func AccountRoutes(route *gin.Engine) {
	route.GET("/account", controllers.FindAccounts)
	route.GET("/account/:id", controllers.FindAccount)
	route.POST("/account", controllers.CreateAccount)
	route.PATCH("/account/:id", controllers.UpdateAccount)
	route.DELETE("/account/:id", controllers.DeleteAccount)
}
