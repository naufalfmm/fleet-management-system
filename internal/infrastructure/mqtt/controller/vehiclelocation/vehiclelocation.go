package vehiclelocation

import (
	"context"
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/naufalfmm/fleet-management-system/internal/model/dto"
	"github.com/naufalfmm/fleet-management-system/internal/pkg/validator"
	"github.com/naufalfmm/fleet-management-system/internal/usecase/vehiclelocation"
)

type Controller struct {
	vehicleLocationUsecase vehiclelocation.Usecase

	validator validator.Validator
}

func NewController(vehicleLocationUsecase vehiclelocation.Usecase, validator validator.Validator) Controller {
	return Controller{
		vehicleLocationUsecase: vehicleLocationUsecase,
		validator:              validator,
	}
}

func (c Controller) InsertLocationFromFleet(cl mqtt.Client, msg mqtt.Message) {
	var payload dto.InsertLocationRequest

	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		return
	}

	if err := c.validator.ValidateStruct(payload); err != nil {
		return
	}

	if err := c.vehicleLocationUsecase.InsertLocationFromFleet(context.Background(), payload); err != nil {
		return
	}
}
