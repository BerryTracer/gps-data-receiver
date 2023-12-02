package model

import "time"

// GPSData represents the data received from a GPS device.
// It includes the device's ID, its latitude and longitude coordinates,
// and the timestamp of when the data was received.
type GPSData struct {
	DeviceID  string    `json:"device_id"` // DeviceID is the unique identifier of the GPS device.
	Latitude  float64   `json:"latitude"`  // Latitude is the geographical latitude coordinate from the GPS device.
	Longitude float64   `json:"longitude"` // Longitude is the geographical longitude coordinate from the GPS device.
	Timestamp time.Time `json:"timestamp"` // Timestamp is the time when the GPS data was received.
	UserID    string    `json:"user_id"`   // UserID is the unique identifier of the user who owns the GPS device.
}
