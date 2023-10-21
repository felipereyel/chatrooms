package broker

import (
	"chatrooms/gosrc/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PostsSubscription interface {
	Close() error
	Channel() <-chan amqp.Delivery
}

type CommandsConsumer interface {
	Close() error
	Channel() <-chan amqp.Delivery
}

type Broker interface {
	Close() error

	PublishPost(roomId string, post models.PostView) error
	SubscribePosts(roomId string) (PostsSubscription, error)

	PublishCommand(command models.CommandView) error
	ConsumeCommands() (CommandsConsumer, error)
}
