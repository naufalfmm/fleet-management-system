package messagequeue

import (
	"context"
)

type MessageQueue interface {
	Publish(ctx context.Context, exchange string, routingKey string, data any) error
	Close() error
}
