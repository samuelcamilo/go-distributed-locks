package controllers

import (
	"example.com/seat-query-api/internal/controllers/seats"
	"example.com/seat-query-api/internal/core/handlers"
)

type (
	Container struct {
		Seat seats.IController
	}
	Options struct {
		Hdl *handlers.Container
	}
)

func New(opts Options) *Container {
	return &Container{
		Seat: seats.New(opts.Hdl),
	}
}
