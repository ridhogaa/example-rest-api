package controller

import (
	"example-rest-api/helper"
	request2 "example-rest-api/model/request"
	"example-rest-api/model/response"
	"example-rest-api/service"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (u *UserControllerImpl) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequest := request2.UserRequest{}
	log.Println("Check request", request)
	helper.ReadFromRequestBody(request, &userRequest)

	userResponse := u.userService.Save(userRequest)
	webResponse := response.BaseResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (u *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usersResponse := u.userService.FindAll()
	webResponse := response.BaseResponse{
		Code:   200,
		Status: "OK",
		Data:   usersResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
