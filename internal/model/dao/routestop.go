package dao

import (
	"time"

	"github.com/shopspring/decimal"
)

type RouteStop struct {
	ID        int64
	RouteID   int64
	Direction string
	Name      string
	Latitude  decimal.Decimal
	Longitude decimal.Decimal
	StopOrder int32
	CreatedAt time.Time
}
