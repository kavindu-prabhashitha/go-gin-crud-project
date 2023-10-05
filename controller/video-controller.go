package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kavindu-prabhashitha/go-gin-project-001/entity"
	"github.com/kavindu-prabhashitha/go-gin-project-001/services"
)

// video controller interface
type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service services.VideoService
}

func New(service services.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
