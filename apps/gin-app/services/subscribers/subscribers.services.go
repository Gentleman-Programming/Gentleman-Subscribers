package services

import (
	"apps/gin-app/models"
	repositoryService "apps/gin-app/services/repository"
	subscriberModel "apps/gin-app/services/subscribers/models"
	"apps/gin-app/utils"

	"go.uber.org/fx"
)

type SubscriberService struct {
	rs *repositoryService.RepositoryService
}

func NewSubscribersService(repositoryService *repositoryService.RepositoryService) SubscriberService {
	return SubscriberService{
		rs: repositoryService,
	}
}

func (s *SubscriberService) ReadAndSetSubscribers() (*[]subscriberModel.Subscriber, *models.AppError) {
	subscribers, err := utils.CsvReader("files/subscribers.csv")
	if err != nil {
		return nil, err
	}

	err = s.rs.SetSubscribers(subscribers)
	if err != nil {
		return nil, err
	}

	return subscribers, nil
}

func (s *SubscriberService) GetAllSubscribers() (*[]subscriberModel.Subscriber, *models.AppError) {
	subscribers, err := s.rs.GetAllSubscribers()

	return subscribers, err
}

var Module = fx.Options(
	fx.Provide(NewSubscribersService),
)
