package listener

import (
	"context"

	"github.com/naufalfmm/fleet-management-system/internal/infrastructure/rabbitmq/controller/vehiclelocation"
	"github.com/naufalfmm/fleet-management-system/internal/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Listener struct {
	vehicleLocationController vehiclelocation.Controller

	logger logger.Logger
}

func NewListener(vehicleLocationController vehiclelocation.Controller, logger logger.Logger) Listener {
	return Listener{
		vehicleLocationController: vehicleLocationController,
		logger:                    logger,
	}
}

func (l *Listener) Listen(conn *amqp.Connection) {
	go l.consumeQueue(conn, "geofence_alerts", l.vehicleLocationController.HandleGeofencingAlert)
}

func (l *Listener) consumeQueue(conn *amqp.Connection, queueName string, handler func(payload []byte)) {
	ch, err := conn.Channel()
	if err != nil {
		l.logger.Error(context.Background(), "failed to open channel", "queue", queueName, "err", err)
		return
	}
	defer ch.Close()

	if _, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		l.logger.Error(context.Background(), "failed to declare queue", "queue", queueName, "err", err)
	}

	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		l.logger.Error(context.Background(), "failed to consume", "queue", queueName, "err", err)
		return
	}

	for msg := range msgs {
		handler(msg.Body)
	}
}
