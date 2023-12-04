package service

import (
	"context"

	"github.com/BerryTracer/gps-data-service/model"
	"github.com/BerryTracer/gps-data-service/repository"
)

type GPSService interface {
	Save(ctx context.Context, gpsData *model.GPSData) error
	FindByDeviceID(ctx context.Context, deviceID string, limit, offset int64) ([]*model.GPSData, error)
	FindByUserID(ctx context.Context, userID string, limit, offset int64) ([]*model.GPSData, error)
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

func (g *GPSServiceImpl) FindByDeviceID(ctx context.Context, deviceID string, limit, offset int64) ([]*model.GPSData, error) {
	return g.Repository.FindByDeviceID(ctx, deviceID, limit, offset)
}

func (g *GPSServiceImpl) FindByUserID(ctx context.Context, userID string, limit, offset int64) ([]*model.GPSData, error) {
	return g.Repository.FindByUserID(ctx, userID, limit, offset)
}

// Ensure GPSServiceImpl implements GPSService
var _ GPSService = (*GPSServiceImpl)(nil)
