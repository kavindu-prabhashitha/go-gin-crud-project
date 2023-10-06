package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kavindu-prabhashitha/go-gin-project-001/entity"
	"github.com/kavindu-prabhashitha/go-gin-project-001/services"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService services.LoginService
	jwtService   services.JWTService
}

func NewLoginController(logService services.LoginService, jwtService services.JWTService) LoginController {
	return &loginController{
		loginService: logService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var reqUser entity.LoginUser

	err := ctx.ShouldBindJSON(&reqUser)
	fmt.Println(reqUser)
	if err != nil {
		return err.Error()
	}
	fmt.Println(reqUser)

	isAuthenticated := controller.loginService.Login(reqUser.UserName, reqUser.Passowrd)
	if isAuthenticated {
		return controller.jwtService.GenerateToken(reqUser.UserName, true)
	}
	return ""
}
