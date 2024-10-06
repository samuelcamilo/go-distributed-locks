package services

import (
	"example.com/ticket-query-api/internal/repositories"
	"example.com/ticket-query-api/internal/services/tickets"
	"example.com/ticket-query-api/pkg/logger"
)

type (
	Container struct {
		Ticket tickets.IService
	}
	Options struct {
		Log  logger.Logger
		Repo *repositories.Container
	}
)

func New(opts Options) *Container {
	return &Container{
		Ticket: tickets.New(opts.Log, opts.Repo),
	}
}
