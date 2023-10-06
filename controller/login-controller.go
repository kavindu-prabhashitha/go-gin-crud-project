package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kavindu-prabhashitha/go-gin-project-001/services"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService services.LoginService
	jwtService   services.JWTService
}

func NewLoginController(logService services.LoginService) LoginController {
	return &loginController{
		loginService: logService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	isAuthenticated := controller.loginService.Login("kavindu", "kp@123")
	if isAuthenticated {
		return controller.jwtService.GenerateToken("kavindu", true)
	}
	return ""
}
