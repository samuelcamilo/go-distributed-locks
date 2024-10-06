package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"

	"example.com/ticket-query-api/internal/repositories/tickets"
	"example.com/ticket-query-api/pkg/logger"
)

type (
	Container struct {
		Ticket tickets.IRepository
	}
	Options struct {
		Log    logger.Logger
		Client *mongo.Client
	}
)

func New(opt Options) *Container {
	return &Container{
		Ticket: tickets.New(opt.Log, opt.Client),
	}
}
