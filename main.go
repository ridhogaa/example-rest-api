package main

import (
	"example-rest-api/app"
	"example-rest-api/controller"
	"example-rest-api/exception"
	"example-rest-api/helper"
	"example-rest-api/middleware"
	"example-rest-api/repository"
	"example-rest-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler

	router.POST("/api/v1/users", userController.Save)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
