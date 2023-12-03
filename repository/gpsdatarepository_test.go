package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/BerryTracer/common-service/adapter/mock"
	"github.com/BerryTracer/gps-data-service/model"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TestMongoGPSDataRepository_Save tests the Save method of the MongoGPSDataRepository
func TestMongoGPSDataRepository_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	repo := NewMongoGPSDataRepository(mockAdapter)

	ctx := context.Background()
	gpsData := &model.GPSData{
		DeviceID:  "test-device",
		Latitude:  12.34,
		Longitude: 56.78,
	}

	// Set expectation on mock
	mockAdapter.EXPECT().
		InsertOne(ctx, gpsData, gomock.Any()).
		Return(&mongo.InsertOneResult{}, nil).
		Times(1)

	err := repo.Save(ctx, gpsData)

	// Assert that the error is nil
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

// TestMongoGPSDataRepository_Save_Error tests the Save method of the MongoGPSDataRepository
func TestMongoGPSDataRepository_Save_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	repo := NewMongoGPSDataRepository(mockAdapter)

	ctx := context.Background()
	gpsData := &model.GPSData{
		DeviceID:  "test-device",
		Latitude:  12.34,
		Longitude: 56.78,
	}

	// Set expectation on mock with an error
	mockAdapter.EXPECT().
		InsertOne(ctx, gpsData, gomock.Any()).
		Return(nil, errors.New("insert error")).
		Times(1)

	err := repo.Save(ctx, gpsData)

	// Assert that there is an error
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

// TestMongoGPSDataRepository_FindByDeviceID tests the FindByDeviceID method of the MongoGPSDataRepository
func TestMongoGPSDataRepository_FindByDeviceID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	repo := NewMongoGPSDataRepository(mockAdapter)

	ctx := context.Background()
	deviceID := "test-device"

	// Mock data
	gpsData := []*model.GPSData{
		{DeviceID: deviceID, Latitude: 12.34, Longitude: 56.78},
		{DeviceID: deviceID, Latitude: 23.45, Longitude: 67.89},
	}

	// Set expectation on mock
	mockCursor := mock.NewMockCursor(ctrl)
	mockAdapter.EXPECT().
		Find(ctx, bson.M{"device_id": deviceID}).
		Return(mockCursor, nil)
	mockCursor.EXPECT().
		All(ctx, gomock.Any()).
		SetArg(1, gpsData).
		Return(nil)
		// In your test function
	mockCursor.EXPECT().Close(gomock.Any()).Return(nil).Times(1)

	// Call the method
	result, err := repo.FindByDeviceID(ctx, deviceID)

	// Assertions
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(result) != len(gpsData) {
		t.Errorf("expected %d results, got %d", len(gpsData), len(result))
	}
}

// TestMongoGPSDataRepository_FindByDeviceID_Error tests the FindByDeviceID method of the MongoGPSDataRepository
func TestMongoGPSDataRepository_FindByDeviceID_AllError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	repo := NewMongoGPSDataRepository(mockAdapter)

	ctx := context.Background()
	deviceID := "test-device"

	mockCursor := mock.NewMockCursor(ctrl)
	mockAdapter.EXPECT().
		Find(ctx, bson.M{"device_id": deviceID}).
		Return(mockCursor, nil)
	mockCursor.EXPECT().
		All(ctx, gomock.Any()).
		Return(errors.New("cursor error"))
	mockCursor.EXPECT().Close(gomock.Any()).Return(nil).Times(1)

	result, err := repo.FindByDeviceID(ctx, deviceID)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if result != nil {
		t.Errorf("expected result to be nil, got %v", result)
	}
}

// TestMongoGPSDataRepository_FindByDeviceID_Error tests the FindByDeviceID method of the MongoGPSDataRepository
func TestMongoGPSDataRepository_FindByDeviceID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	repo := NewMongoGPSDataRepository(mockAdapter)

	ctx := context.Background()
	deviceID := "test-device"

	mockAdapter.EXPECT().
		Find(ctx, bson.M{"device_id": deviceID}).
		Return(nil, errors.New("find error"))

	result, err := repo.FindByDeviceID(ctx, deviceID)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if result != nil {
		t.Errorf("expected result to be nil, got %v", result)
	}
}

// TestMongoGPSDataRepository_FindByUserID tests the FindByUserID method of the MongoGPSDataRepository
func TestMongoGPSDataRepository_FindByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	repo := NewMongoGPSDataRepository(mockAdapter)

	ctx := context.Background()
	userID := "test-user"

	// Mock data
	gpsData := []*model.GPSData{
		{UserID: userID, Latitude: 12.34, Longitude: 56.78},
		{UserID: userID, Latitude: 23.45, Longitude: 67.89},
	}

	// Set expectation on mock
	mockCursor := mock.NewMockCursor(ctrl)
	mockAdapter.EXPECT().
		Find(ctx, bson.M{"user_id": userID}).
		Return(mockCursor, nil)
	mockCursor.EXPECT().
		All(ctx, gomock.Any()).
		SetArg(1, gpsData).
		Return(nil)
		// In your test function
	mockCursor.EXPECT().Close(gomock.Any()).Return(nil).Times(1)

	// Call the method
	result, err := repo.FindByUserID(ctx, userID)

	// Assertions
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(result) != len(gpsData) {
		t.Errorf("expected %d results, got %d", len(gpsData), len(result))
	}
}

// TestMongoGPSDataRepository_FindByUserID_Error tests the FindByUserID method of the MongoGPSDataRepository
func TestMongoGPSDataRepository_FindByUserID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	repo := NewMongoGPSDataRepository(mockAdapter)

	ctx := context.Background()
	userID := "test-user"

	mockAdapter.EXPECT().
		Find(ctx, bson.M{"user_id": userID}).
		Return(nil, errors.New("find error"))

	result, err := repo.FindByUserID(ctx, userID)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if result != nil {
		t.Errorf("expected result to be nil, got %v", result)
	}
}

// TestMongoGPSDataRepository_FindByDeviceID_AllError tests the FindByDeviceID method of the MongoGPSDataRepository
func TestMongoGPSDataRepository_FindByUserID_AllError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdapter := mock.NewMockMongoAdapter(ctrl)
	repo := NewMongoGPSDataRepository(mockAdapter)

	ctx := context.Background()
	userID := "test-user"

	mockCursor := mock.NewMockCursor(ctrl)
	mockAdapter.EXPECT().
		Find(ctx, bson.M{"user_id": userID}).
		Return(mockCursor, nil)
	mockCursor.EXPECT().
		All(ctx, gomock.Any()).
		Return(errors.New("cursor error"))
	mockCursor.EXPECT().Close(gomock.Any()).Return(nil).Times(1)

	result, err := repo.FindByUserID(ctx, userID)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if result != nil {
		t.Errorf("expected result to be nil, got %v", result)
	}
}
