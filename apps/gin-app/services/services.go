package services

import (
	repositoryService "apps/gin-app/services/repository"
	subscriberService "apps/gin-app/services/subscribers"

	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(repositoryService.Module, subscriberService.Module)
