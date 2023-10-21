package broker

import (
	"chatrooms/gosrc/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type subscription interface {
	Close() error
	Channel() <-chan amqp.Delivery
}

type Broker interface {
	Close() error
	Publish(roomId string, post models.PostView) error
	Subscribe(roomId string) (subscription, error)
}
