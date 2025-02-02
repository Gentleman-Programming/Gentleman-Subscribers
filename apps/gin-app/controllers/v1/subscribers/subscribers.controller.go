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

func (ctrl *SubscriberController) ReadAndSetSubscribers(c *gin.Context) {
	subscribers, err := ctrl.service.ReadAndSetSubscribers()
	if err != nil {
		c.JSON(err.Code, utils.ResponseError{Message: err.Message})
	}
	c.JSON(http.StatusOK, subscribers)
}

func (ctrl *SubscriberController) GetSubscribrers(c *gin.Context) {
	subscribers, err := ctrl.service.GetAllSubscribers()
	if err != nil {
		c.JSON(err.Code, utils.ResponseError{Message: err.Message})
	}

	c.JSON(http.StatusOK, subscribers)
}
