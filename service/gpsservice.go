package service

import (
	"context"

	"github.com/BerryTracer/gps-data-service/repository"

	"github.com/BerryTracer/gps-data-service/model"
)

type GPSService interface {
	Save(ctx context.Context, gpsData *model.GPSData) error
	FindByDeviceID(ctx context.Context, deviceID string) ([]*model.GPSData, error)
	FindByUserID(ctx context.Context, userID string) ([]*model.GPSData, error)
}

type GPSServiceImpl struct {
	Repository repository.GPSDataRepository
}

func NewGPSService(repository repository.GPSDataRepository) *GPSServiceImpl {
	return &GPSServiceImpl{Repository: repository}
}

func (g *GPSServiceImpl) Save(ctx context.Context, gpsData *model.GPSData) error {
	return g.Repository.Save(ctx, gpsData)
}

func (g *GPSServiceImpl) FindByDeviceID(ctx context.Context, deviceID string) ([]*model.GPSData, error) {
	return g.Repository.FindByDeviceID(ctx, deviceID)
}

func (g *GPSServiceImpl) FindByUserID(ctx context.Context, userID string) ([]*model.GPSData, error) {
	return g.Repository.FindByUserID(ctx, userID)
}

// Ensure GPSServiceImpl implements GPSService
var _ GPSService = (*GPSServiceImpl)(nil)
