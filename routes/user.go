package routes

import (
	controller "mytodoapp/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}
