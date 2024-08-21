package seats

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"example.com/seat-query-api/internal/core/models"
	"example.com/seat-query-api/pkg/logger"
)

type (
	IRepository interface {
		GetById(id int) (seat models.SeatModel, err error)
	}
	seatRepo struct {
		log    logger.Logger
		client *mongo.Client
	}
)

func New(log logger.Logger, client *mongo.Client) IRepository {
	return &seatRepo{
		log:    log,
		client: client,
	}
}

func (repo *seatRepo) GetById(id int) (seat models.SeatModel, err error) {
	coll := repo.client.Database("sessions-db").Collection("SESSIONS")

	var result models.SeatModel

	filter := bson.D{primitive.E{Key: "session_id", Value: id}}
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		repo.log.Error(err.Error())
		return models.SeatModel{}, err
	}

	return result, nil
}
