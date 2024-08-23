package handlers

import (
	"example.com/event-query-api/internal/core/handlers/events"
	"example.com/event-query-api/internal/services"
	"example.com/event-query-api/pkg/logger"
)

type (
	Container struct {
		Event events.IHandler
	}
	Options struct {
		Log logger.Logger
		Srv *services.Container
	}
)

func New(opts Options) *Container {
	return &Container{
		Event: events.New(opts.Log, opts.Srv),
	}
}
