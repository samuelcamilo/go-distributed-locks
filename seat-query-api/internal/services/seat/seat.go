package seat

import (
	"example.com/seat-query-api/internal/core/models"
	"example.com/seat-query-api/pkg/logger"
)

type (
	IService interface {
		GetAll() (seats []models.SeatModel, err error)
	}

	services struct {
		log logger.Logger
	}
)

func New(log logger.Logger) IService {
	return &services{log: log}
}

func (s *services) GetAll() (seats []models.SeatModel, err error) {
	seats = []models.SeatModel{
		{
			Id:        1,
			SessionId: 21323,
			Elements: []models.Element{
				{
					Row:         1,
					Col:         0,
					Code:        "O 1",
					Name:        "Seat O 1",
					Description: "Cadeira Lovers",
					Status:      1,
				},
			},
		},
	}

	return
}
