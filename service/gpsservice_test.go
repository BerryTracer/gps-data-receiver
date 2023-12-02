package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/BerryTracer/gps-data-service/repository/mock"

	"github.com/BerryTracer/gps-data-service/model"

	"github.com/golang/mock/gomock"
)

func TestGPSServiceImpl_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockGPSDataRepository(ctrl)
	service := NewGPSService(mockRepo)

	ctx := context.Background()
	gpsData := &model.GPSData{ /* ... set test data ... */ }

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

	mockRepo := mock.NewMockGPSDataRepository(ctrl)
	service := NewGPSService(mockRepo)

	ctx := context.Background()
	deviceID := "test-device"

	// Mock data
	gpsData := []*model.GPSData{ /* ... set test data ... */ }

	// Success scenario
	mockRepo.EXPECT().FindByDeviceID(ctx, deviceID).Return(gpsData, nil)
	result, err := service.FindByDeviceID(ctx, deviceID)
	if err != nil {
		t.Errorf("FindByDeviceID() error = %v, wantErr %v", err, nil)
	}
	if !reflect.DeepEqual(result, gpsData) {
		t.Errorf("FindByDeviceID() = %v, want %v", result, gpsData)
	}

	// Error scenario
	mockRepo.EXPECT().FindByDeviceID(ctx, deviceID).Return(nil, errors.New("find error"))
	_, err = service.FindByDeviceID(ctx, deviceID)
	if err == nil {
		t.Error("FindByDeviceID() expected an error, got nil")
	}
}

func TestGPSServiceImpl_FindByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockGPSDataRepository(ctrl)
	service := NewGPSService(mockRepo)

	ctx := context.Background()
	userID := "test-user"

	// Mock data
	gpsData := []*model.GPSData{ /* ... set test data ... */ }

	// Success scenario
	mockRepo.EXPECT().FindByUserID(ctx, userID).Return(gpsData, nil)
	result, err := service.FindByUserID(ctx, userID)
	if err != nil {
		t.Errorf("FindByUserID() error = %v, wantErr %v", err, nil)
	}
	if !reflect.DeepEqual(result, gpsData) {
		t.Errorf("FindByUserID() = %v, want %v", result, gpsData)
	}

	// Error scenario
	mockRepo.EXPECT().FindByUserID(ctx, userID).Return(nil, errors.New("find error"))
	_, err = service.FindByUserID(ctx, userID)
	if err == nil {
		t.Error("FindByUserID() expected an error, got nil")
	}
}
