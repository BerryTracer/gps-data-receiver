package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GPSDatabase struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewGPSDatabaseConnection(connStr string) (*GPSDatabase, error) {
	clientOptions := options.Client().ApplyURI(connStr).SetMaxPoolSize(50)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	db := client.Database("gps_data")
	collection := db.Collection("gps_data")
	return &GPSDatabase{Client: client, Collection: collection}, nil
}

func (d *GPSDatabase) Disconnect() error {
	return d.Client.Disconnect(context.Background())
}
