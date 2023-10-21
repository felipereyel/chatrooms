package broker

import (
	"chatrooms/gosrc/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type postsSubscription interface {
	Close() error
	Channel() <-chan amqp.Delivery
}

type commandsConsumer interface {
	Close() error
	Channel() <-chan amqp.Delivery
}

type Broker interface {
	Close() error

	PublishPost(roomId string, post models.PostView) error
	SubscribePosts(roomId string) (postsSubscription, error)

	PublishCommand(command models.CommandView) error
	ConsumeCommands() (commandsConsumer, error)
}
