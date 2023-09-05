package router

import (
	"mytodoapp/controller"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(UsersController *controller.UsersController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	usersRouter := baseRouter.Group("/users")
	usersRouter.GET("", UsersController.FindALL)
	usersRouter.GET("/:userId", UsersController.FindById)
	usersRouter.POST("", UsersController.Create)
	usersRouter.PATCH("/:userId", UsersController.Update)
	usersRouter.DELETE("/:userId", UsersController.Delete)

	return router
}
