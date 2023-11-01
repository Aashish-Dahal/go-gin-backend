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

type CommentController struct {
	logger         infrastructure.Logger
	commentService services.CommentService
	env            infrastructure.Env
}

func NewCommentController(logger infrastructure.Logger,
	commentService services.CommentService,
	env infrastructure.Env) CommentController {
	return CommentController{
		logger:         logger,
		commentService: commentService,
		env:            env,
	}
}
func (cc CommentController) CreateComment(c *gin.Context) {
	comment := models.Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		cc.logger.Zap.Error("Error [CreateComment] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind Comment")
		responses.HandleError(c, err)
		return
	}
	if err := cc.commentService.Create(comment); err != nil {
		cc.logger.Zap.Error("Error [CreateComment] [db CreateComment]: ", err.Error())
		err := errors.BadRequest.Wrap(err, "Failed To Create Comment")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Comment Created Successfully")

}
