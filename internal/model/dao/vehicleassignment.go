package dao

import "time"

type VehicleAssignment struct {
	ID                 int64
	VehicleID          string
	AssignedRouteID    *int64
	CurrentRouteStopID *int64
	AssignedDirection  *string
	CreatedAt          time.Time
}
