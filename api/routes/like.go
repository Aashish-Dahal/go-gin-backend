package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

type LikeRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	likeController controllers.LikeController
	trxMiddleware  middlewares.DBTransactionMiddleware
}

func NewLikeRoutes(logger infrastructure.Logger,
	router infrastructure.Router,
	likeController controllers.LikeController,
	trxMiddleware middlewares.DBTransactionMiddleware) LikeRoutes {
	return LikeRoutes{
		logger:         logger,
		router:         router,
		likeController: likeController,
		trxMiddleware:  trxMiddleware,
	}

}

// Setup user routes
func (i LikeRoutes) Setup() {
	i.logger.Zap.Info(" Setting up feed routes")
	likes := i.router.Gin.Group("/like")
	{
		likes.POST("", i.trxMiddleware.DBTransactionHandle(), i.likeController.CreateLike)

	}
}
