package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

type CommentRoutes struct {
	logger            infrastructure.Logger
	router            infrastructure.Router
	commentController controllers.CommentController
	trxMiddleware     middlewares.DBTransactionMiddleware
}

func NewCommentRoutes(logger infrastructure.Logger,
	router infrastructure.Router,
	commentController controllers.CommentController,
	trxMiddleware middlewares.DBTransactionMiddleware) CommentRoutes {
	return CommentRoutes{
		logger:            logger,
		router:            router,
		commentController: commentController,
		trxMiddleware:     trxMiddleware,
	}

}

// Setup user routes
func (i CommentRoutes) Setup() {
	i.logger.Zap.Info(" Setting up feed routes")
	comments := i.router.Gin.Group("/comment")
	{
		comments.POST("", i.trxMiddleware.DBTransactionHandle(), i.commentController.CreateComment)

	}
}
