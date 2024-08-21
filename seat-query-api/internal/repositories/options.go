package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"

	"example.com/seat-query-api/internal/repositories/seats"
	"example.com/seat-query-api/pkg/logger"
)

type (
	Container struct {
		Seat seats.IRepository
	}
	Options struct {
		Log    logger.Logger
		Client *mongo.Client
	}
)

func New(opt Options) *Container {
	return &Container{
		Seat: seats.New(opt.Log, opt.Client),
	}
}
