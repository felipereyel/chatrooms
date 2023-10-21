package broker

import (
	"chatrooms/gosrc/models"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type fakeBroker struct {
	postsChan    map[string]chan amqp.Delivery
	commandsChan chan amqp.Delivery
}

func FakeBrokerRepo() Broker {
	postsChan := make(map[string]chan amqp.Delivery)
	commandsChan := make(chan amqp.Delivery)

	return &fakeBroker{postsChan, commandsChan}
}

func (b *fakeBroker) Close() error {
	close(b.commandsChan)
	for _, ch := range b.postsChan {
		close(ch)
	}
	return nil
}

func (b *fakeBroker) PublishPost(roomId string, post models.PostView) error {
	body, err := json.Marshal(post)
	if err != nil {
		return err
	}

	ch, ok := b.postsChan[roomId]
	if !ok {
		ch = make(chan amqp.Delivery)
		b.postsChan[roomId] = ch
	}

	ch <- amqp.Delivery{Body: body}
	return nil
}

type fakeSubscription struct {
	channel <-chan amqp.Delivery
}

func (c fakeSubscription) Close() error {
	return nil
}

func (c fakeSubscription) Channel() <-chan amqp.Delivery {
	return c.channel
}

func (b *fakeBroker) SubscribePosts(roomId string) (postsSubscription, error) {
	ch, ok := b.postsChan[roomId]
	if !ok {
		ch = make(chan amqp.Delivery)
		b.postsChan[roomId] = ch
	}

	return &fakeSubscription{ch}, nil
}

func (b *fakeBroker) PublishCommand(command models.CommandView) error {
	body, err := json.Marshal(command)
	if err != nil {
		return err
	}

	b.commandsChan <- amqp.Delivery{Body: body}
	return nil
}

type fakeConsumer struct {
	channel <-chan amqp.Delivery
}

func (c fakeConsumer) Close() error {
	return nil
}

func (c fakeConsumer) Channel() <-chan amqp.Delivery {
	return c.channel
}

func (b *fakeBroker) ConsumeCommands() (commandsConsumer, error) {
	return &fakeConsumer{b.commandsChan}, nil
}
