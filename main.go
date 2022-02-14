package main

import (
	"github.com/gin-gonic/gin"
	"github.com/seyi-js/go-rest-api/controllers"
)

func main() {
	r := gin.Default()

	// models.ConnectDatabase()

	r.GET("/", controllers.FindBooks)

	r.Run()
}
