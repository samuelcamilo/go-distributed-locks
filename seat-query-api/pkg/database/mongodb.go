package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MustConnect(s string) *mongo.Client {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(s),
	)
	if err != nil {
		panic(err)
	}

	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()

	return client
}
