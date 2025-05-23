package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	Connection *amqp.Connection
}

func (e *Emitter) Setup() error {
	channel, err := e.Connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	return declareExchange(channel)

}

func (e *Emitter) Push(event, severity string) error {
	channel, err := e.Connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	log.Println("Pushing to channel")

	err = channel.Publish(
		"logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	if err != nil {
		return err
	}
	return nil

}

func NewEventEmitter(conn *amqp.Connection) (*Emitter, error) {
	emitter := &Emitter{
		Connection: conn,
	}
	err := emitter.Setup()
	if err != nil {
		return nil, err
	}
	return emitter, nil
}
