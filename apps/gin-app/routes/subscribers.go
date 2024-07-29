package routes

import (
	subscribers_v1 "apps/gin-app/controllers/v1/subscribers"
	"apps/gin-app/utils"
)

// MiscRoutes struct
type SubscribersRoutes struct {
	subscriberController subscribers_v1.SubscriberController
	logger               utils.Logger
	handler              utils.RequestHandler
}

// Setup Misc routes
func (s SubscribersRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/apis/v1")
	{
		api.GET("/getAllSubscribers", s.subscriberController.GetSubscribrers)
		api.GET("/setAllSubscribers", s.subscriberController.ReadAndSetSubscribers)
	}
}

// NewMiscRoutes creates new Misc controller
func NewSubscriberRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	subscriberController subscribers_v1.SubscriberController,
) SubscribersRoutes {
	return SubscribersRoutes{
		handler:              handler,
		logger:               logger,
		subscriberController: subscriberController,
	}
}
