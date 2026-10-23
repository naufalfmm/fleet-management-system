package dto

import "github.com/shopspring/decimal"

type InsertLocationRequest struct {
	VehicleID string `json:"-"`

	Latitude  decimal.Decimal `json:"latitude" validate:"dreq,dgte=-5.320,dlte=-6.398"`
	Longitude decimal.Decimal `json:"longitude" validate:"dreq,dgte=106.378,dlte=107.017"`
	Timestamp int64           `json:"timestamp" validate:"required"`
}

type GetVehicleLocationHistoriesRequest struct {
	VehicleID string `form:"-"`

	TimestampFrom int64 `form:"start"`
	TimestampTo   int64 `form:"end"`
}
