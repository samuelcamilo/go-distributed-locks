package events

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"example.com/event-query-api/internal/core/models"
	"example.com/event-query-api/pkg/logger"
)

type (
	IRepository interface {
		GetAll(ctx context.Context) (events []models.EventModel, err error)
		GetById(id int) (event *models.EventModel, err error)
	}
	eventRepo struct {
		log    logger.Logger
		client *mongo.Client
	}
)

func New(log logger.Logger, client *mongo.Client) IRepository {
	return &eventRepo{
		log:    log,
		client: client,
	}
}

func (repo *eventRepo) GetAll(ctx context.Context) (events []models.EventModel, err error) {
	coll := repo.client.Database("events-db").Collection("EVENTS")

	cursor, err := coll.Find(ctx, bson.D{})
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return
}

func (repo *eventRepo) GetById(id int) (event *models.EventModel, err error) {
	coll := repo.client.Database("events-db").Collection("EVENTS")

	filter := bson.D{primitive.E{Key: "event_id", Value: id}}
	err = coll.FindOne(context.TODO(), filter).Decode(&event)

	return
}
