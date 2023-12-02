package api

import (
	"github.com/BerryTracer/gps-data-service/service"

	"github.com/BerryTracer/gps-data-service/model"
)

type GPSHandler struct {
	Service service.GPSService
}

func NewGPSHandler(service service.GPSService) *GPSHandler {
	return &GPSHandler{Service: service}
}

func (g *GPSHandler) SaveGPSData(c HttpContext) error {
	var gpsData model.GPSData
	err := c.Bind(&gpsData)
	if err != nil {
		c.JSON(400, err)
		return err
	}

	return g.Service.Save(c.Context(), &gpsData)
}
