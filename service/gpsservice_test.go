package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/BerryTracer/gps-data-service/model"
	mock_repository "github.com/BerryTracer/gps-data-service/repository/mock"

	"github.com/golang/mock/gomock"
)

func TestGPSServiceImpl_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockGPSDataRepository(ctrl)
	service := NewGPSService(mockRepo)

	ctx := context.Background()
	gpsData := &model.GPSData{
		DeviceID:  "1",
		Latitude:  37.7749,
		Longitude: -122.4194,
		Timestamp: time.Now(),
		UserID:    "test-user",
	}

	// Success scenario
	mockRepo.EXPECT().Save(ctx, gpsData).Return(nil)
	err := service.Save(ctx, gpsData)
	if err != nil {
		t.Errorf("Save() error = %v, wantErr %v", err, nil)
	}

	// Error scenario
	mockRepo.EXPECT().Save(ctx, gpsData).Return(errors.New("save error"))
	err = service.Save(ctx, gpsData)
	if err == nil {
		t.Error("Save() expected an error, got nil")
	}
}

func TestGPSServiceImpl_FindByDeviceID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockGPSDataRepository(ctrl)
	service := NewGPSService(mockRepo)

	ctx := context.Background()
	deviceID := "test-device"
	limit := int64(10)
	offset := int64(0)

	// Mock data
	gpsData := []*model.GPSData{
		{
			DeviceID:  "1",
			Latitude:  37.7749,
			Longitude: -122.4194,
			Timestamp: time.Now(),
			UserID:    "test-user",
		},
		{
			DeviceID:  "2",
			Latitude:  34.0522,
			Longitude: -118.2437,
			Timestamp: time.Now(),
			UserID:    "test-user",
		},
	}

	// Success scenario
	mockRepo.EXPECT().FindByDeviceID(ctx, deviceID, limit, offset).Return(gpsData, nil)
	result, err := service.FindByDeviceID(ctx, deviceID, limit, offset)
	if err != nil {
		t.Errorf("FindByDeviceID() error = %v, wantErr %v", err, nil)
	}
	if !reflect.DeepEqual(result, gpsData) {
		t.Errorf("FindByDeviceID() = %v, want %v", result, gpsData)
	}

	// Error scenario
	mockRepo.EXPECT().FindByDeviceID(ctx, deviceID, limit, offset).Return(nil, errors.New("find error"))
	_, err = service.FindByDeviceID(ctx, deviceID, limit, offset)
	if err == nil {
		t.Error("FindByDeviceID() expected an error, got nil")
	}
}

func TestGPSServiceImpl_FindByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockGPSDataRepository(ctrl)
	service := NewGPSService(mockRepo)

	ctx := context.Background()
	userID := "test-user"
	limit := int64(10)
	offset := int64(0)

	// Mock data
	gpsData := []*model.GPSData{
		{
			DeviceID:  "1",
			Latitude:  37.7749,
			Longitude: -122.4194,
			Timestamp: time.Now(),
			UserID:    "test-user",
		},
		{
			DeviceID:  "2",
			Latitude:  34.0522,
			Longitude: -118.2437,
			Timestamp: time.Now(),
			UserID:    "test-user",
		},
	}

	// Success scenario
	mockRepo.EXPECT().FindByUserID(ctx, userID, limit, offset).Return(gpsData, nil)
	result, err := service.FindByUserID(ctx, userID, limit, offset)
	if err != nil {
		t.Errorf("FindByUserID() error = %v, wantErr %v", err, nil)
	}
	if !reflect.DeepEqual(result, gpsData) {
		t.Errorf("FindByUserID() = %v, want %v", result, gpsData)
	}

	// Error scenario
	mockRepo.EXPECT().FindByUserID(ctx, userID, limit, offset).Return(nil, errors.New("find error"))
	_, err = service.FindByUserID(ctx, userID, limit, offset)
	if err == nil {
		t.Error("FindByUserID() expected an error, got nil")
	}
}
