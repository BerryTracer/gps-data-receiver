package repository

import (
	"context"

	"github.com/BerryTracer/gps-data-service/model"

	"github.com/BerryTracer/common-service/adapter"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GPSDataRepository interface {
	Save(ctx context.Context, gpsData *model.GPSData) error
	FindByDeviceID(ctx context.Context, deviceID string) ([]*model.GPSData, error)
	FindByUserID(ctx context.Context, userID string) ([]*model.GPSData, error)
}

type MongoGPSDataRepository struct {
	Collection adapter.MongoAdapter
}

func NewMongoGPSDataRepository(collection adapter.MongoAdapter) *MongoGPSDataRepository {
	return &MongoGPSDataRepository{Collection: collection}
}

func (m *MongoGPSDataRepository) Save(ctx context.Context, gpsData *model.GPSData) error {
	_, err := m.Collection.InsertOne(ctx, gpsData)
	return err
}

func (m *MongoGPSDataRepository) FindByDeviceID(ctx context.Context, deviceID string) ([]*model.GPSData, error) {
	filter := primitive.M{"device_id": deviceID}
	cursor, err := m.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var gpsData []*model.GPSData
	if err := cursor.All(ctx, &gpsData); err != nil {
		return nil, err
	}

	return gpsData, nil
}

func (m *MongoGPSDataRepository) FindByUserID(ctx context.Context, userID string) ([]*model.GPSData, error) {
	filter := primitive.M{"user_id": userID}
	cursor, err := m.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var gpsData []*model.GPSData
	if err := cursor.All(ctx, &gpsData); err != nil {
		return nil, err
	}

	return gpsData, nil
}

// Ensure MongoGPSDataRepository implements GPSDataRepository
var _ GPSDataRepository = (*MongoGPSDataRepository)(nil)
