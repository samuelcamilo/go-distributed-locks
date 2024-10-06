package tickets

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"example.com/ticket-query-api/internal/core/models"
	"example.com/ticket-query-api/pkg/logger"
)

type (
	IRepository interface {
		GetAll(ctx context.Context) (tickets []models.TicketModel, err error)
		GetById(ctx context.Context, id int) (ticket *models.TicketModel, err error)
	}
	repositories struct {
		log    logger.Logger
		client *mongo.Client
	}
)

func New(log logger.Logger, client *mongo.Client) IRepository {
	return &repositories{
		log:    log,
		client: client,
	}
}

func (r *repositories) GetAll(ctx context.Context) (tickets []models.TicketModel, err error) {
	coll := r.client.Database("andromeda").Collection("tickets")

	cursor, err := coll.Find(ctx, bson.D{})
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &tickets); err != nil {
		return nil, err
	}

	return
}

func (r *repositories) GetById(
	ctx context.Context,
	id int,
) (ticket *models.TicketModel, err error) {
	coll := r.client.Database("andromeda").Collection("tickets")

	filter := bson.D{primitive.E{Key: "ticket_id", Value: id}}
	err = coll.FindOne(ctx, filter).Decode(&ticket)

	return
}
