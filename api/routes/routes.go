package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewDocsRoutes),
	fx.Provide(NewJwtAuthRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewFeedRoutes),
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
	jwtAuthRoutes JwtAuthRoutes,
	docsRoutes DocsRoutes,
) Routes {
	return Routes{

		jwtAuthRoutes,
		userRoutes,
		docsRoutes,
		feedRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
