package controller

import (
	"mytodoapp/data1/request"
	"mytodoapp/data1/response"
	"mytodoapp/helper"
	"mytodoapp/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UsersController struct {
	usersService service.UsersService
}

func NewUsersController(service service.UsersService) *UsersController {
	return &UsersController{
		usersService: service,
	}
}

func (controller *UsersController) Create(ctx *gin.Context) {
	log.Info().Msg("create users")
	CreateUsersRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&CreateUsersRequest)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UsersController) Update(ctx *gin.Context) {
	log.Info().Msg("Update users")
	updateUsersRequest := request.UpdateUsersRequest{}
	err := ctx.ShouldBindJSON(&updateUsersRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	updateUsersRequest.Id = id

	controller.usersService.Update(updateUsersRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UsersController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete Users")
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controller.usersService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/Json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UsersController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid users")
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	userResponse := controller.usersService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UsersController) FindALL(ctx *gin.Context) {
	log.Info().Msg("findAll users")
	userResponse := controller.usersService.FindALL()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
