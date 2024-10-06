package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketModel struct {
	Id        primitive.ObjectID `bson:"_id"        json:"_id"`
	TicketId  int                `bson:"ticket_id"  json:"ticket_id"`
	StatusId  int                `bson:"status_id"  json:"status_id"`
	Status    string             `bson:"status"     json:"status"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time         `bson:"updated_at" json:"updated_at"`
}
