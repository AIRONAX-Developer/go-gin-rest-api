package controller

import (
	"fmt"
	locred "github.com/AIRONAX-Developer/go-gin-rest-api/src/LoCred"
	service "github.com/AIRONAX-Developer/go-gin-rest-api/src/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type LoginControllerStruct struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func LoginHandler(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &LoginControllerStruct{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *LoginControllerStruct) Login(ctx *gin.Context) string {
	var credentials locred.LoginCredentialsStruct
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		fmt.Println("ctx.ShouldBind ==>> ", err)
		return err.Error()
	}

	isUserAuthenticated := controller.loginService.LoginUser(credentials.Email, credentials.Password)
	if isUserAuthenticated {
		return controller.jwtService.GenerateToken(credentials.Email, true)
	}
	return ""
}
