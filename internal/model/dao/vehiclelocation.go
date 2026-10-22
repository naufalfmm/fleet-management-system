package dao

import (
	"time"

	"github.com/shopspring/decimal"
)

type VehicleLocation struct {
	ID        int64
	VehicleID string
	Latitude  decimal.Decimal
	Longitude decimal.Decimal
	Timestamp int64
	CreatedAt time.Time
}
