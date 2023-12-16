package api

import (
	"errors"
	"strconv"

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
		if jsonErr := c.JSON(400, err); jsonErr != nil {
			return jsonErr
		}
		return err
	}

	err = gpsData.Validate()
	if err != nil {
		if jsonErr := c.JSON(400, err); jsonErr != nil {
			return jsonErr
		}
		return err
	}

	return g.Service.Save(c.Context(), &gpsData)
}

func (g *GPSHandler) GetGPSDataByDeviceId(c HttpContext) error {
	deviceId := c.Params("device_id")
	if deviceId == "" {
		if jsonErr := c.JSON(400, errors.New("device_id is required")); jsonErr != nil {
			return jsonErr
		}
		return errors.New("device_id is required")
	}

	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		offset = 0
	}

	gpsDataArray, err := g.Service.FindByDeviceID(c.Context(), deviceId, limit, offset)
	if err != nil {
		if jsonErr := c.JSON(400, err); jsonErr != nil {
			return jsonErr
		}
		return err
	}

	return c.JSON(200, gpsDataArray)
}

func (g *GPSHandler) GetGPSDataByUserId(c HttpContext) error {
	userId := c.Params("user_id")
	if userId == "" {
		if jsonErr := c.JSON(400, errors.New("user_id is required")); jsonErr != nil {
			return jsonErr
		}
		return errors.New("user_id is required")
	}

	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		offset = 0
	}

	gpsDataArray, err := g.Service.FindByUserID(c.Context(), userId, limit, offset)
	if err != nil {
		if jsonErr := c.JSON(400, err); jsonErr != nil {
			return jsonErr
		}
		return err
	}

	return c.JSON(200, gpsDataArray)
}
