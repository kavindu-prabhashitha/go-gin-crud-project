package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kavindu-prabhashitha/go-gin-project-001/entity"
	"github.com/kavindu-prabhashitha/go-gin-project-001/services"
	"github.com/kavindu-prabhashitha/go-gin-project-001/validators"
)

// video controller interface
type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) (entity.Video, error)
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service services.VideoService
}

var validate *validator.Validate

func New(service services.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return entity.Video{}, err
	}
	err = validate.Struct(video)
	if err != nil {
		return entity.Video{}, err
	}
	c.service.Save(video)
	return video, err
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
