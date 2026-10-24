package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/naufalfmm/fleet-management-system/internal/pkg/messagequeue"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ(amqpUrl string) (messagequeue.MessageQueue, error) {
	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &rabbitMQ{
		conn:    conn,
		channel: ch,
	}, nil
}

func (r *rabbitMQ) Publish(ctx context.Context, exchange string, routingKey string, data any) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = r.channel.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = r.channel.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *rabbitMQ) Close() error {
	if r.channel != nil {
		return r.channel.Close()
	}

	if r.conn != nil {
		return r.conn.Close()
	}

	return nil
}
