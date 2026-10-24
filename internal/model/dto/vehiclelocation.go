package dto

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/fleet-management-system/internal/common"
	"github.com/shopspring/decimal"
)

type InsertLocationRequest struct {
	VehicleID string          `json:"vehicle_id" validate:"required"`
	Latitude  decimal.Decimal `json:"latitude" validate:"dreq"`
	Longitude decimal.Decimal `json:"longitude" validate:"dreq"`
	Timestamp int64           `json:"timestamp" validate:"required"`
}

type GetVehicleLocationHistoriesRequest struct {
	VehicleID     string
	TimestampFrom int64
	TimestampTo   int64
}

func (r *GetVehicleLocationHistoriesRequest) FromGinContext(gc *gin.Context) error {
	vehicleID := gc.Param("vehicle_id")
	if vehicleID == "" {
		return common.ErrMissingParam("vehicle_id")
	}

	strTimestampFrom := gc.Query("start")
	if strTimestampFrom != "" {
		timestampFrom, err := strconv.ParseInt(strTimestampFrom, 10, 64)
		if err == nil {
			r.TimestampFrom = timestampFrom
		}
	}

	strTimestampTo := gc.Query("start")
	if strTimestampTo != "" {
		timestampTo, err := strconv.ParseInt(strTimestampTo, 10, 64)
		if err == nil {
			r.TimestampTo = timestampTo
		}
	}

	return nil
}

type VehicleLocationResponse struct {
	VehicleID string          `json:"vehicle_id"`
	Latitude  decimal.Decimal `json:"latitude"`
	Longitude decimal.Decimal `json:"longitude"`
	Timestamp int64           `json:"timestamp"`
}
