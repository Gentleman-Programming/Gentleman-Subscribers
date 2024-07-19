package subscribers_v1

import (
	services "apps/gin-app/services/subscribers"
	"apps/gin-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubscriberController struct {
	service services.SubscriberService
}

func NewSubscriberController(service services.SubscriberService) SubscriberController {
	return SubscriberController{service: service}
}

func (ctrl *SubscriberController) GetSubscribrers(c *gin.Context) {
	subscribers, err := ctrl.service.GetAllSubscribers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError{Message: "Error getting all Subscribers"})
	}

	c.JSON(http.StatusOK, subscribers)
}
