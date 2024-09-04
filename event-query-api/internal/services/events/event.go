package events

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"example.com/event-query-api/internal/core/models"
	"example.com/event-query-api/internal/repositories"
	"example.com/event-query-api/pkg/logger"
)

type (
	IService interface {
		GetAll() (events []models.EventModel, err error)
		GetById(id int) (event *models.EventModel, err error)
	}

	eventService struct {
		log  logger.Logger
		repo *repositories.Container
	}
)

func New(log logger.Logger, repo *repositories.Container) IService {
	return &eventService{log: log, repo: repo}
}

func (s *eventService) GetAll() (events []models.EventModel, err error) {
	events, err = s.repo.Event.GetAll(context.TODO())
	if err != nil {
		return nil, err
	}

	if len(events) < 1 {
		return []models.EventModel{}, nil
	}

	return
}

func (s *eventService) GetById(id int) (event *models.EventModel, err error) {
	event, err = s.repo.Event.GetById(id)
	if err == mongo.ErrNoDocuments {
		return nil, ErrorEventNotFound
	}

	return
}
