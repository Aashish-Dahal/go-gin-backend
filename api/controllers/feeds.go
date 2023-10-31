package controllers

import (
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/paginations"
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
func (cc FeedController) GetAllFeeds(c *gin.Context) {
	pagination := paginations.BuildPagination[*paginations.FeedPagination](c)

	feeds, count, err := cc.feedService.GetAllFeeds(*pagination)
	if err != nil {
		cc.logger.Zap.Error("Error finding feed records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get feeds data")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, feeds, count)
}
func (cc FeedController) GetFeedByID(c *gin.Context) {
	pagination := paginations.BuildPagination[*paginations.FeedPagination](c)

	idParam := c.Param("id")

	feeds, count, err := cc.feedService.GetFeedByID(*pagination,idParam)
	if err != nil {
		cc.logger.Zap.Error("Error finding user feeds", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get users feed data")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, feeds, count)
}
func (cc FeedController) DeleteFeed(c *gin.Context) {

	idParam := c.Param("id")

	 err := cc.feedService.DeleteFeed(idParam)
	if err != nil {
		cc.logger.Zap.Error("Error finding feeds", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to delete feed data")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Feed deleted Successfully")
}