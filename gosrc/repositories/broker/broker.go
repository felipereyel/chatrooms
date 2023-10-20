package broker

import (
	"chatrooms/gosrc/config"
	"chatrooms/gosrc/models"
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// type PubBroker interface {
// 	Close() error
// 	Publish(roomId string, post models.PostView) error
// }

// type SubBroker interface {
// 	Subscribe(roomId string) (<-chan models.PostView, error)
// }

type broker struct {
	conn *amqp.Connection
}

const exchangeName = "chatrooms"

func NewBrokerRepo() (*broker, error) {
	conn, err := amqp.Dial(config.Configs.RabbitMQConnString)
	if err != nil {
		return nil, err
	}

	return &broker{conn}, nil
}

func (b *broker) Close() error {
	return b.conn.Close()
}

func (b *broker) Publish(roomId string, post models.PostView) error {
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
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return ch.PublishWithContext(ctx,
		exchangeName, // exchange
		roomId,       // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (c *subscription) Close() error {
	return c.brokerChannel.Close()
}

func (b *broker) Subscribe(roomId string) (*subscription, error) {
	ch, err := b.conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
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
		q.Name,       // queue name
		roomId,       // routing key
		exchangeName, // exchange
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

	return &subscription{ch, msgs}, nil
}
