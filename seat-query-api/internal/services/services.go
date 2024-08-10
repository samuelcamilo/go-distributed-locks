package services

import (
	"example.com/seat-query-api/internal/services/seat"
	"example.com/seat-query-api/pkg/logger"
)

type (
	Container struct {
		Seat seat.IService
	}

	Options struct {
		Log logger.Logger
	}
)

func New(opts Options) *Container {
	return &Container{
		Seat: seat.New(opts.Log),
	}
}
