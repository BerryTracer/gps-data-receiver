package repository

import (
	"context"

	"github.com/BerryTracer/common-service/adapter/database/mongodb"
	"github.com/BerryTracer/gps-data-service/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GPSDataRepository interface {
	Save(ctx context.Context, gpsData *model.GPSData) error
	FindByDeviceID(ctx context.Context, deviceID string, page, pageSize int64) ([]*model.GPSData, error)
	FindByUserID(ctx context.Context, userID string, page, pageSize int64) ([]*model.GPSData, error)
}

type MongoGPSDataRepository struct {
	Collection mongodb.MongoAdapter
}

func NewMongoGPSDataRepository(collection mongodb.MongoAdapter) *MongoGPSDataRepository {
	return &MongoGPSDataRepository{Collection: collection}
}

func (m *MongoGPSDataRepository) Save(ctx context.Context, gpsData *model.GPSData) error {
	_, err := m.Collection.InsertOne(ctx, gpsData)
	return err
}

func (m *MongoGPSDataRepository) FindByDeviceID(ctx context.Context, deviceID string, page, pageSize int64) ([]*model.GPSData, error) {
	filter := primitive.M{"device_id": deviceID}
	skip := (page - 1) * pageSize
	options := options.Find().SetSkip(skip).SetLimit(pageSize)

	cursor, err := m.Collection.Find(ctx, filter, options)
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

func (m *MongoGPSDataRepository) FindByUserID(ctx context.Context, userID string, page, pageSize int64) ([]*model.GPSData, error) {
	filter := primitive.M{"user_id": userID}
	skip := (page - 1) * pageSize
	options := options.Find().SetSkip(skip).SetLimit(pageSize)

	cursor, err := m.Collection.Find(ctx, filter, options)
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
