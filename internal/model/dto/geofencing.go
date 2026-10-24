package dto

import "github.com/shopspring/decimal"

type GeofenceEventMessageLocation struct {
	Latitude  decimal.Decimal `json:"latitude" validate:"dreq"`
	Longitude decimal.Decimal `json:"longitude" validate:"dreq"`
}

type GeofenceEventMessage struct {
	VehicleID string                       `json:"vehicle_id" validate:"required"`
	Event     string                       `json:"event" validate:"required"`
	Location  GeofenceEventMessageLocation `json:"location" validate:"required"`
	Timestamp int64                        `json:"timestamp" validate:"required"`
}
