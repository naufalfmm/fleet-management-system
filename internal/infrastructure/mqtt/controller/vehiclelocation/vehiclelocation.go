package vehiclelocation

import (
	"context"
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/naufalfmm/fleet-management-system/internal/model/dto"
	"github.com/naufalfmm/fleet-management-system/internal/pkg/logger"
	"github.com/naufalfmm/fleet-management-system/internal/pkg/validator"
	"github.com/naufalfmm/fleet-management-system/internal/usecase/vehiclelocation"
)

type Controller struct {
	vehicleLocationUsecase vehiclelocation.Usecase

	validator validator.Validator
	logger    logger.Logger
}

func NewController(vehicleLocationUsecase vehiclelocation.Usecase, validator validator.Validator, logger logger.Logger) Controller {
	return Controller{
		vehicleLocationUsecase: vehicleLocationUsecase,
		validator:              validator,
		logger:                 logger,
	}
}

func (c Controller) InsertLocationFromFleet(cl mqtt.Client, msg mqtt.Message) {
	ctx := context.Background()

	var payload dto.InsertLocationRequest
	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		c.logger.Error(ctx, "error when parsing message", "err", err)
		return
	}

	if err := c.validator.ValidateStruct(payload); err != nil {
		c.logger.Error(ctx, "error format data", "err", err)
		return
	}

	vehicleLoc, err := c.vehicleLocationUsecase.InsertLocationFromFleet(ctx, payload)
	if err != nil {
		c.logger.Error(ctx, "error insert location", "err", err)
		return
	}

	vehicleLocation := dto.VehicleLocationResponse{
		VehicleID: vehicleLoc.VehicleID,
		Latitude:  vehicleLoc.Latitude,
		Longitude: vehicleLoc.Longitude,
		Timestamp: vehicleLoc.Timestamp,
	}

	c.logger.Info(ctx, "location is inserted successfully", "data", vehicleLocation)
}
