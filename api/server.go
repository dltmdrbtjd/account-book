package main

import (
	"api/middlewares"
	"api/models"
	"api/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.JSONMiddleware())

	routers.AccountRoutes(r)
	if err := r.Run(":8081"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
