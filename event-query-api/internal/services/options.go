package services

import (
	"example.com/event-query-api/internal/repositories"
	"example.com/event-query-api/internal/services/events"
	"example.com/event-query-api/pkg/logger"
)

type (
	Container struct {
		Event events.IService
	}

	Options struct {
		Log  logger.Logger
		Repo *repositories.Container
	}
)

func New(opts Options) *Container {
	return &Container{
		Event: events.New(opts.Log, opts.Repo),
	}
}
