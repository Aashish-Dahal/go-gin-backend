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

type FeedController struct {
	logger      infrastructure.Logger
	feedService services.FeedService
	env         infrastructure.Env
}

func NewFeedController(logger infrastructure.Logger,
	feedService services.FeedService,
	env infrastructure.Env) FeedController {
	return FeedController{
		logger:      logger,
		feedService: feedService,
		env:         env,
	}
}
func (cc FeedController) CreateFeed(c *gin.Context) {
	feed := models.Feed{}
	if err := c.ShouldBindJSON(&feed); err != nil {
		cc.logger.Zap.Error("Error [CreateFeed] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind Feed")
		responses.HandleError(c, err)
		return
	}
	if err := cc.feedService.CreateFeed(feed); err != nil {
		cc.logger.Zap.Error("Error [CreateFeed] [db CreateFeed]: ", err.Error())
		err := errors.BadRequest.Wrap(err, "Failed To Create Feed")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Feed Created Successfully")

}
