package subscriber

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/naufalfmm/fleet-management-system/config"
	"github.com/naufalfmm/fleet-management-system/internal/infrastructure/mqtt/controller/vehiclelocation"
)

type Subscriber struct {
	vehicleLocationController vehiclelocation.Controller

	conf config.Config
}

func NewSubscriber(vehicleLocationController vehiclelocation.Controller, conf config.Config) Subscriber {
	return Subscriber{
		vehicleLocationController: vehicleLocationController,
		conf:                      conf,
	}
}

func (s Subscriber) Register(client mqtt.Client) {
	client.Subscribe("fleet/vehicle/+/location", byte(s.conf.RabbitMqSubscriberQos), s.vehicleLocationController.InsertLocationFromFleet).Wait()
}
