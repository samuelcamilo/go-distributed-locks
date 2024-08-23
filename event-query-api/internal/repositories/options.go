package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"

	"example.com/event-query-api/internal/repositories/events"
	"example.com/event-query-api/pkg/logger"
)

type (
	Container struct {
		Event events.IRepository
	}
	Options struct {
		Log    logger.Logger
		Client *mongo.Client
	}
)

func New(opt Options) *Container {
	return &Container{
		Event: events.New(opt.Log, opt.Client),
	}
}
