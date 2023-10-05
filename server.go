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
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	apiRoutes := server.Group("/api")
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

	server.Run(":8080")
}
