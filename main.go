package main

import (
	"example-rest-api/app"
	"example-rest-api/controller"
	"example-rest-api/exception"
	"example-rest-api/helper"
	"example-rest-api/repository"
	"example-rest-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := app.NewDB()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)
	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler

	router.POST("/api/v1/users", userController.Save)
	router.GET("/api/v1/users", userController.FindAll)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
