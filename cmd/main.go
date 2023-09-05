package main

import (
	"mytodoapp/controller"
	"mytodoapp/database"
	"mytodoapp/helper"
	"mytodoapp/models"
	"mytodoapp/repository"
	"mytodoapp/router"
	"mytodoapp/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server!")

	//Database
	db := database.DatabaseConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&models.Users{})

	//Repository
	usersRepository := repository.NewUsersREpositoryImpl(db)

	//service
	usersService := service.NewUsersServiceImpl(usersRepository, validate)

	//Controller
	usersController := controller.NewUsersController(usersService)

	//Router
	routes := router.NewRouter(usersController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
