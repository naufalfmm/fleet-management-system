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

func (v VehicleLocation) Args() []any {
	return []any{
		"vehicle_id", v.VehicleID,
		"latitude", v.Latitude,
		"longitude", v.Longitude,
		"timestamp", v.Timestamp,
	}
}
