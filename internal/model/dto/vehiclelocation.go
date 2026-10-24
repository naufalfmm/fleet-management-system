package dto

import "github.com/shopspring/decimal"

type InsertLocationRequest struct {
	VehicleID string          `json:"vehicle_id" validate:"required"`
	Latitude  decimal.Decimal `json:"latitude" validate:"dreq"`
	Longitude decimal.Decimal `json:"longitude" validate:"dreq"`
	Timestamp int64           `json:"timestamp" validate:"required"`
}

type GetVehicleLocationHistoriesRequest struct {
	VehicleID string `form:"-"`

	TimestampFrom int64 `form:"start"`
	TimestampTo   int64 `form:"end"`
}
