package controller

import (
	"github.com/BayronCampaz/gin-rest-api/entity"
	"github.com/BayronCampaz/gin-rest-api/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type videoController struct {
	service service.VideoService
}

func New(s service.VideoService) VideoController {
	return &videoController{
		service: s,
	}
}

func (vc *videoController) FindAll() []entity.Video {
	return vc.service.FindAll()
}

func (vc *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	vc.service.Save(video)
	return video
}
