package model

import (
	"time"

	grpc "github.com/BerryTracer/gps-data-service/grpc/proto"
	"github.com/go-playground/validator/v10"
)

// GPSData represents the data received from a GPS device.
// It includes the device's ID, its latitude and longitude coordinates,
// and the timestamp of when the data was received.
type GPSData struct {
	DeviceID  string    `bson:"device_id" json:"device_id" validate:"required"`
	Latitude  float64   `bson:"latitude" json:"latitude" validate:"latitude"`
	Longitude float64   `bson:"longitude" json:"longitude" validate:"longitude"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp" validate:"required"`
	UserID    string    `bson:"user_id" json:"user_id" validate:"required"`
}

// Validate validates the GPSData object.
func (gpsData *GPSData) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("latitude", validateLatitude)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("longitude", validateLongitude)
	if err != nil {
		return err
	}
	return validate.Struct(gpsData)
}

// validateLatitude validates the latitude coordinate.
func validateLatitude(fl validator.FieldLevel) bool {
	latitude := fl.Field().Float()
	return latitude >= -90 && latitude <= 90
}

// validateLongitude validates the longitude coordinate.
func validateLongitude(fl validator.FieldLevel) bool {
	longitude := fl.Field().Float()
	return longitude >= -180 && longitude <= 180
}

// ConvertToProto converts a GPSData object to a proto.GPSData object.
func (gpsData *GPSData) ConvertToProto() *grpc.GPSData {
	return &grpc.GPSData{
		DeviceId:  gpsData.DeviceID,
		Latitude:  gpsData.Latitude,
		Longitude: gpsData.Longitude,
		Timestamp: gpsData.Timestamp.Unix(),
		UserId:    gpsData.UserID,
	}
}
