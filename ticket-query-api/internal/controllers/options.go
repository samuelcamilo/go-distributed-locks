package controllers

import (
	"example.com/ticket-query-api/internal/controllers/tickets"
	"example.com/ticket-query-api/internal/core/handlers"
)

type (
	Container struct {
		Ticket tickets.IController
	}
	Options struct {
		Hdl *handlers.Container
	}
)

func New(opts Options) *Container {
	return &Container{
		Ticket: tickets.New(opts.Hdl),
	}
}
