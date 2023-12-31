package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewDocsRoutes),
	fx.Provide(NewJwtAuthRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewFeedRoutes),
	fx.Provide(NewLikeRoutes),
	fx.Provide(NewCommentRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	userRoutes UserRoutes,
	feedRoutes FeedRoutes,
	likeRoutes LikeRoutes,
	commentRoutes CommentRoutes,
	jwtAuthRoutes JwtAuthRoutes,
	docsRoutes DocsRoutes,
) Routes {
	return Routes{

		jwtAuthRoutes,
		userRoutes,
		docsRoutes,
		feedRoutes,
		likeRoutes,
		commentRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
