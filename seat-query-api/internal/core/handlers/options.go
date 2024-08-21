package handlers

import (
	"example.com/seat-query-api/internal/core/handlers/seats"
	"example.com/seat-query-api/internal/services"
	"example.com/seat-query-api/pkg/logger"
)

type (
	Container struct {
		Seat seats.IHandler
	}
	Options struct {
		Log logger.Logger
		Srv *services.Container
	}
)

func New(opts Options) *Container {
	return &Container{
		Seat: seats.New(opts.Log, opts.Srv),
	}
}
