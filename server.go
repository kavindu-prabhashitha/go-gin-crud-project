package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kavindu-prabhashitha/go-gin-project-001/controller"
	"github.com/kavindu-prabhashitha/go-gin-project-001/middlewares"
	"github.com/kavindu-prabhashitha/go-gin-project-001/services"
)

var (
	videoService    services.VideoService      = services.New()
	VideoController controller.VideoController = controller.New(videoService)

	loginService    services.LoginService      = services.NewLoginService()
	loginController controller.LoginController = controller.NewLoginController(loginService)
)

func setupLogOutput() {
	f, err := os.Create("gin.log")
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.Println("Init Gin log set ")
}

func main() {
	setupLogOutput()
	server := gin.New()
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")
	server.Use(gin.Recovery(), middlewares.Logger())

	// Login Endpoint : Authentication + Token Creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT(), middlewares.Logger())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, VideoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			video, err := VideoController.Save(ctx)
			if err != nil {
				fmt.Println(err)
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(200, video)
			}

		})

	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", VideoController.ShowAll)
	}

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "OK!!",
	// 	})
	// })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
