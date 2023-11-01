package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

type FeedRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	feedController controllers.FeedController
	trxMiddleware  middlewares.DBTransactionMiddleware
}

func NewFeedRoutes(logger infrastructure.Logger,
	router infrastructure.Router,
	feedController controllers.FeedController,
	trxMiddleware middlewares.DBTransactionMiddleware) FeedRoutes {
	return FeedRoutes{
		logger:         logger,
		router:         router,
		feedController: feedController,
		trxMiddleware:  trxMiddleware,
	}

}

// Setup user routes
func (i FeedRoutes) Setup() {
	i.logger.Zap.Info(" Setting up feed routes")
	feeds := i.router.Gin.Group("/feed")
	{
		feeds.POST("", i.trxMiddleware.DBTransactionHandle(), i.feedController.CreateFeed)
		feeds.GET("", i.feedController.GetAllFeeds)
		feeds.GET("/:id", i.feedController.GetFeedByID)
		feeds.DELETE("/:id", i.feedController.DeleteFeed)

	}
}
