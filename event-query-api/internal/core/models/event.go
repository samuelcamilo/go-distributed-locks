package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventModel struct{}

type (
	SeatModel struct {
		Id        primitive.ObjectID `bson:"_id"`
		SessionId int32              `bson:"session_id,omitempty" json:"session_id"`
		Elements  []Element          `bson:"seats,omitempty"      json:"seats"`
		CreatedAt *time.Time         `bson:"created_at"           json:"created_at"`
	}
	Element struct {
		Id     int    `bson:"id"     json:"id"`
		Line   int    `bson:"line"   json:"line"`
		Column int    `bson:"column" json:"column"`
		Label  string `bson:"label"  json:"label"`
		Type   string `bson:"type"   json:"type"`
		Status string `bson:"status" json:"status"`
	}
)
