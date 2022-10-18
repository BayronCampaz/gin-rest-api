package controller

import (
	"net/http"

	"github.com/BayronCampaz/gin-rest-api/entity"
	"github.com/BayronCampaz/gin-rest-api/service"
	"github.com/BayronCampaz/gin-rest-api/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type videoController struct {
	service service.VideoService
}

var validate *validator.Validate

func New(s service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: s,
	}
}

func (vc *videoController) FindAll() []entity.Video {
	return vc.service.FindAll()
}

func (vc *videoController) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	vc.service.Save(video)
	return nil
}

func (vc *videoController) ShowAll(ctx *gin.Context) {
	videos := vc.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
