package tickets

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"example.com/ticket-query-api/internal/core/models"
	"example.com/ticket-query-api/internal/repositories"
	"example.com/ticket-query-api/pkg/logger"
)

type (
	IService interface {
		GetAll(ctx context.Context) (tickets []models.TicketModel, err error)
		GetById(ctx context.Context, id int) (ticket *models.TicketModel, err error)
	}
	services struct {
		log  logger.Logger
		repo *repositories.Container
	}
)

func New(log logger.Logger, repo *repositories.Container) IService {
	return &services{log: log, repo: repo}
}

func (s *services) GetAll(ctx context.Context) (tickets []models.TicketModel, err error) {
	tickets, err = s.repo.Ticket.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(tickets) < 1 {
		return []models.TicketModel{}, nil
	}

	return
}

func (s *services) GetById(
	ctx context.Context,
	id int,
) (ticket *models.TicketModel, err error) {
	ticket, err = s.repo.Ticket.GetById(ctx, id)
	if err == mongo.ErrNoDocuments {
		fmt.Println(err)
		return nil, ErrorTicketNotFound
	}

	return
}
