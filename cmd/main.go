package main

import (
	"mytodoapp/database"
	"mytodoapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello World!")
	// })
	database.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}
