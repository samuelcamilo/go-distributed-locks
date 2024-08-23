package controllers

import (
	"example.com/event-query-api/internal/controllers/events"
	"example.com/event-query-api/internal/core/handlers"
)

type (
	Container struct {
		Event events.IController
	}
	Options struct {
		Hdl *handlers.Container
	}
)

func New(opts Options) *Container {
	return &Container{
		Event: events.New(opts.Hdl),
	}
}
