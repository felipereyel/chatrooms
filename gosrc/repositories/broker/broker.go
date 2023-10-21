package broker

import (
	"chatrooms/gosrc/config"
	"chatrooms/gosrc/models"
	"context"
	"encoding/json"
	"errors"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type broker struct {
	conn *amqp.Connection
}

const postsExchangeName = "chatrooms:posts"
const commandsQueueName = "chatrooms:commands"

func NewBrokerRepo() (*broker, error) {
	if config.Configs.RabbitMQConnString == "" {
		return nil, errors.New("RabbitMQConnString is not set")
	}

	conn, err := amqp.Dial(config.Configs.RabbitMQConnString)
	if err != nil {
		return nil, err
	}

	return &broker{conn}, nil
}

func (b *broker) Close() error {
	return b.conn.Close()
}

func (b *broker) PublishPost(roomId string, post models.PostView) error {
	body, err := json.Marshal(post)
	if err != nil {
		return err
	}

	ch, err := b.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		postsExchangeName, // name
		"topic",           // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return ch.PublishWithContext(ctx,
		postsExchangeName, // exchange
		roomId,            // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

type rabitSubscription struct {
	brokerChannel *amqp.Channel
	MessageChan   <-chan amqp.Delivery
}

func (c rabitSubscription) Close() error {
	return c.brokerChannel.Close()
}

func (c rabitSubscription) Channel() <-chan amqp.Delivery {
	return c.MessageChan
}

func (b *broker) SubscribePosts(roomId string) (postsSubscription, error) {
	ch, err := b.conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		postsExchangeName, // name
		"topic",           // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,            // queue name
		roomId,            // routing key
		postsExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return nil, err
	}

	return &rabitSubscription{ch, msgs}, nil
}

func (b *broker) PublishCommand(command models.CommandView) error {
	body, err := json.Marshal(command)
	if err != nil {
		return err
	}

	ch, err := b.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		commandsQueueName, // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(body),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

type rabitConsumer struct {
	brokerChannel *amqp.Channel
	MessageChan   <-chan amqp.Delivery
}

func (c rabitConsumer) Close() error {
	return c.brokerChannel.Close()
}

func (c rabitConsumer) Channel() <-chan amqp.Delivery {
	return c.MessageChan
}

func (b *broker) ConsumeCommands() (commandsConsumer, error) {
	ch, err := b.conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		commandsQueueName, // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return nil, err
	}

	return &rabitConsumer{ch, msgs}, nil
}
