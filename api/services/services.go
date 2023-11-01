package services

import "go.uber.org/fx"

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewFirebaseService),
	fx.Provide(NewUserService),
	fx.Provide(NewJWTAuthService),
	fx.Provide(NewFeedService),
	fx.Provide(NewLikeService),
	fx.Provide(NewCommentService),
)
