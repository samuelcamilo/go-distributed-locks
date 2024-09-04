package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddressModel struct {
	Street     string `bson:"street"      json:"street"`
	City       string `bson:"city"        json:"city"`
	State      string `bson:"state"       json:"state"`
	PostalCode string `bson:"postal_code" json:"postal_code"`
	Country    string `bson:"country"     json:"country"`
}

type EventModel struct {
	Id               primitive.ObjectID `bson:"_id"`
	EventId          int                `bson:"event_id"          json:"event_id"`
	Name             string             `bson:"name"              json:"name"`
	Description      string             `bson:"description"       json:"description"`
	Category         string             `bson:"category"          json:"category"`
	StartAt          time.Time          `bson:"start_at"          json:"start_at"`
	EndAt            time.Time          `bson:"end_at"            json:"end_at"`
	StatusId         int                `bson:"status_id"         json:"status_id"`
	Status           string             `bson:"status"            json:"status"`
	AvailableTickets int                `bson:"available_tickets" json:"available_tickets"`
	Address          *AddressModel      `bson:"address"           json:"address"`
	ImageUrl         string             `bson:"image_url"         json:"image_url"`
	CreatedAt        time.Time          `bson:"created_at"        json:"created_at"`
	UpdatedAt        *time.Time         `bson:"updated_at"        json:"updated_at"`
}
