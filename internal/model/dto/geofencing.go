package dto

import "github.com/shopspring/decimal"

type GeofenceEventMessageLocation struct {
	Latitude  decimal.Decimal `json:"latitude"`
	Longitude decimal.Decimal `json:"longitude"`
}

type GeofenceEventMessage struct {
	VehicleID string                       `json:"vehicle_id"`
	Event     string                       `json:"event"`
	Location  GeofenceEventMessageLocation `json:"location"`
	Timestamp int64                        `json:"timestamp"`
}
