package handlers

import (
	"example.com/ticket-query-api/internal/core/handlers/tickets"
	"example.com/ticket-query-api/internal/services"
	"example.com/ticket-query-api/pkg/logger"
)

type (
	Container struct {
		Ticket tickets.IHandler
	}
	Options struct {
		Log logger.Logger
		Srv *services.Container
	}
)

func New(opts Options) *Container {
	return &Container{
		Ticket: tickets.New(opts.Log, opts.Srv),
	}
}
