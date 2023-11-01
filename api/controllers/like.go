package controllers

import (
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	logger      infrastructure.Logger
	likeService services.LikeService
	env         infrastructure.Env
}

func NewLikeController(logger infrastructure.Logger,
	likeService services.LikeService,
	env infrastructure.Env) LikeController {
	return LikeController{
		logger:      logger,
		likeService: likeService,
		env:         env,
	}
}
func (cc LikeController) CreateLike(c *gin.Context) {
	like := models.Like{}
	if err := c.ShouldBindJSON(&like); err != nil {
		cc.logger.Zap.Error("Error [CreateLike] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind Like")
		responses.HandleError(c, err)
		return
	}
	if err := cc.likeService.Create(like); err != nil {
		cc.logger.Zap.Error("Error [CreateLike] [db CreateLike]: ", err.Error())
		err := errors.BadRequest.Wrap(err, "Failed To Create Like")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Like Created Successfully")

}
