package model

import (
	"time"

	grpc "github.com/BerryTracer/gps-data-service/grpc/proto"
)

// GPSData represents the data received from a GPS device.
// It includes the device's ID, its latitude and longitude coordinates,
// and the timestamp of when the data was received.
type GPSData struct {
	DeviceID  string    `bson:"device_id" json:"device_id"`
	Latitude  float64   `bson:"latitude" json:"latitude"`
	Longitude float64   `bson:"longitude" json:"longitude"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	UserID    string    `bson:"user_id" json:"user_id"`
}

func (gpsData *GPSData) ConvertToProto() *grpc.GPSData {
	return &grpc.GPSData{
		DeviceId:  gpsData.DeviceID,
		Latitude:  gpsData.Latitude,
		Longitude: gpsData.Longitude,
		Timestamp: gpsData.Timestamp.Unix(),
		UserId:    gpsData.UserID,
	}
}
