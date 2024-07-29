package services

import (
	subscribersService "apps/gin-app/services/subscribers"

	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	subscribersService.Module,
)
