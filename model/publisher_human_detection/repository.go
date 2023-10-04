package publisher_human_detection

import (
	"context"
	"encoding/json"

	"github.com/aditya3232/gatewatchApp-services.git/connection"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Repository interface {
	CreateQueueHumanDetection(input HumanDetectionInput) (HumanDetection, error)
}

type repository struct {
	rabbitmq *amqp.Connection
}

func NewRepository(rabbitmq *amqp.Connection) *repository {
	return &repository{rabbitmq}
}

func (r *repository) CreateQueueHumanDetection(input HumanDetectionInput) (HumanDetection, error) {
	ch, err := connection.RabbitMQ().Channel()
	if err != nil {
		return HumanDetection{}, err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"HumanDetectionQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return HumanDetection{}, err
	}

	// body from input
	var inputReadytoMarshal = HumanDetectionInput{
		Tid:           input.Tid,
		DateTime:      input.DateTime,
		Person:        input.Person,
		ConvertedFile: input.ConvertedFile,
	}

	// Convert the HumanDetection struct to JSON
	body, err := json.Marshal(inputReadytoMarshal)
	if err != nil {
		return HumanDetection{}, err
	}

	ctx := context.Background() // Create a context
	err = ch.PublishWithContext(ctx,
		"",
		"HumanDetectionQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return HumanDetection{}, err
	}

	// Return the sent HumanDetection struct
	return HumanDetection{}, err
}
