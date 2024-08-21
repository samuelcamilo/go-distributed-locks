package services

import (
	"example.com/seat-query-api/internal/repositories"
	"example.com/seat-query-api/internal/services/seat"
	"example.com/seat-query-api/pkg/logger"
)

type (
	Container struct {
		Seat seats.IService
	}

	Options struct {
		Log  logger.Logger
		Repo *repositories.Container
	}
)

func New(opts Options) *Container {
	return &Container{
		Seat: seats.New(opts.Log, opts.Repo),
	}
}
