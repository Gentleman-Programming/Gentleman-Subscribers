package controllers

import (
	subscribers_v1 "apps/gin-app/controllers/v1/subscribers"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(subscribers_v1.NewSubscriberController),
)
