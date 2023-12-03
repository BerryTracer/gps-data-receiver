package api

import (
	"errors"

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

func (g *GPSHandler) GetGPSDataByDeviceId(c HttpContext) error {
	deviceId := c.Query("device_id")
	if deviceId == "" {
		c.SendStatus(400)
		return errors.New("device_id is required")
	}

	gpsDataArray, err := g.Service.FindByDeviceID(c.Context(), deviceId)
	if err != nil {
		c.JSON(400, err)
		return err
	}

	c.JSON(200, gpsDataArray)
	return nil
}
