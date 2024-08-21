package seats

import (
	"example.com/seat-query-api/internal/core/models"
	"example.com/seat-query-api/internal/repositories"
	"example.com/seat-query-api/pkg/logger"
)

type (
	IService interface {
		GetById(id int) (seat models.SeatModel, err error)
	}

	services struct {
		log  logger.Logger
		repo *repositories.Container
	}
)

func New(log logger.Logger, repo *repositories.Container) IService {
	return &services{log: log, repo: repo}
}

func (s *services) GetById(id int) (seat models.SeatModel, err error) {
	seat, err = s.repo.Seat.GetById(id)
	if err != nil {
		return seat, nil
	}
	return
}
