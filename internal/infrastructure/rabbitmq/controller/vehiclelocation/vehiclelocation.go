package vehiclelocation

import (
	"context"
	"encoding/json"

	"github.com/naufalfmm/fleet-management-system/internal/model/dto"
	pkgerror "github.com/naufalfmm/fleet-management-system/internal/pkg/error"
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

func (c Controller) HandleGeofencingAlert(payload []byte) {
	var message dto.GeofenceEventMessage
	if err := json.Unmarshal(payload, &message); err != nil {
		c.logger.Error(context.Background(), "error unmarshal payload", "err", err)
		return
	}

	if err := c.validator.ValidateStruct(message); err != nil {
		c.logger.Error(context.Background(), "validate error", "err", pkgerror.GenerateError(message, err))
	}

	c.logger.Info(context.Background(), "handle geofencing alert is successful", "message", message)
}
