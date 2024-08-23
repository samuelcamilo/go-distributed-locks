package events

import (
	"go.mongodb.org/mongo-driver/mongo"

	"example.com/event-query-api/internal/core/models"
	"example.com/event-query-api/internal/repositories"
	"example.com/event-query-api/pkg/logger"
)

type (
	IService interface {
		GetById(id int) (event *models.SeatModel, err error)
	}

	services struct {
		log  logger.Logger
		repo *repositories.Container
	}
)

func New(log logger.Logger, repo *repositories.Container) IService {
	return &services{log: log, repo: repo}
}

func (s *services) GetById(id int) (event *models.SeatModel, err error) {
	event, err = s.repo.Event.GetById(id)
	if err == mongo.ErrNoDocuments {
		return nil, ErrorEventNotFound
	}

	return
}
