package vehiclelocation

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/fleet-management-system/internal/common"
	"github.com/naufalfmm/fleet-management-system/internal/model/dto"
	"github.com/naufalfmm/fleet-management-system/internal/usecase/vehiclelocation"
)

type Controller struct {
	vehicleLocationUsecase vehiclelocation.Usecase
}

func NewController(vehicleLocationUsecase vehiclelocation.Usecase) Controller {
	return Controller{
		vehicleLocationUsecase: vehicleLocationUsecase,
	}
}

func (c Controller) GetLatestLocation(gc *gin.Context) {
	vehicleID := gc.Param("vehicle_id")
	if vehicleID == "" {
		err := common.ErrMissingParam("vehicle_id")
		gc.JSON(http.StatusBadRequest, dto.Response[dto.ErrorResponse]{
			Ok:      false,
			Code:    common.ErrorCodeMissingParam,
			Message: err.Error(),
			Data: dto.ErrorResponse{
				Error: err,
			},
		})

		return
	}

	latestLoc, err := c.vehicleLocationUsecase.GetLatestVehicleLocation(gc.Request.Context(), vehicleID)
	if err != nil {
		c.buildErrorForGetLatestLocation(gc, err)
		return
	}

	gc.JSON(http.StatusOK, dto.Response[dto.VehicleLocationResponse]{
		Ok:      true,
		Code:    common.SuccessCodeOK,
		Message: "Success",
		Data: dto.VehicleLocationResponse{
			VehicleID: latestLoc.VehicleID,
			Latitude:  latestLoc.Latitude,
			Longitude: latestLoc.Longitude,
			Timestamp: latestLoc.Timestamp,
		},
	})
}

func (c Controller) buildErrorForGetLatestLocation(gc *gin.Context, err error) {
	statusCode := http.StatusInternalServerError

	if errors.Is(err, common.ErrNoLocationCmp) {
		gc.JSON(http.StatusForbidden, dto.Response[dto.ErrorResponse]{
			Ok:      false,
			Code:    common.ErrorCodeNoLocation,
			Message: err.Error(),
			Data: dto.ErrorResponse{
				Error: err,
			},
		})
		return
	}

	gc.JSON(statusCode, dto.Response[dto.ErrorResponse]{
		Ok:      false,
		Code:    common.ErrorCodeInternalError,
		Message: common.ErrDefaultError.Error(),
		Data: dto.ErrorResponse{
			Error: common.ErrDefaultError,
		},
	})
}

func (c Controller) GetLocationHistories(gc *gin.Context) {
	var req dto.GetVehicleLocationHistoriesRequest
	if err := req.FromGinContext(gc); err != nil {
		c.buildErrorForGetLocationHistories(gc, err)
		return
	}

	locationHistories, err := c.vehicleLocationUsecase.GetVehicleLocationHistories(gc.Request.Context(), req)
	if err != nil {
		c.buildErrorForGetLocationHistories(gc, err)
		return
	}

	locationHistoryResps := make([]dto.VehicleLocationResponse, len(locationHistories))
	for i, locationHistory := range locationHistories {
		locationHistoryResps[i] = dto.VehicleLocationResponse{
			VehicleID: locationHistory.VehicleID,
			Latitude:  locationHistory.Latitude,
			Longitude: locationHistory.Longitude,
			Timestamp: locationHistory.Timestamp,
		}
	}

	gc.JSON(http.StatusOK, dto.Response[[]dto.VehicleLocationResponse]{
		Ok:      true,
		Code:    common.SuccessCodeOK,
		Message: "Success",
		Data:    locationHistoryResps,
	})
}

func (c Controller) buildErrorForGetLocationHistories(gc *gin.Context, err error) {
	statusCode := http.StatusInternalServerError

	if errors.Is(err, common.ErrMissingParamCmp) {
		gc.JSON(http.StatusForbidden, dto.Response[dto.ErrorResponse]{
			Ok:      false,
			Code:    common.ErrorCodeMissingParam,
			Message: err.Error(),
			Data: dto.ErrorResponse{
				Error: err,
			},
		})
		return
	}

	gc.JSON(statusCode, dto.Response[dto.ErrorResponse]{
		Ok:      false,
		Code:    common.ErrorCodeInternalError,
		Message: common.ErrDefaultError.Error(),
		Data: dto.ErrorResponse{
			Error: common.ErrDefaultError,
		},
	})
}
