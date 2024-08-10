package handlers

import (
	"example.com/seat-query-api/internal/core/handlers/seat"
	"example.com/seat-query-api/internal/services"
	"example.com/seat-query-api/pkg/logger"
)

type (
	Container struct {
		Seat seat.IHandler
	}
	Options struct {
		Srv *services.Container
		Log logger.Logger
	}
)

func New(opts Options) *Container {
	return &Container{
		Seat: seat.New(opts.Srv, opts.Log),
	}
}
