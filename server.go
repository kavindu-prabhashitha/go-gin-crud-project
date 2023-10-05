package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kavindu-prabhashitha/go-gin-project-001/controller"
	"github.com/kavindu-prabhashitha/go-gin-project-001/services"
)

var (
	videoService    services.VideoService      = services.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":8080")
}
