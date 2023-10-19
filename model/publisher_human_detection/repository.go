package publisher_human_detection

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Repository interface {
	/*
		- ambil input dari API, dimasukkan ke entity rmq
		- disini parameter adalah entitiy
		- disini returnnya yg akan ditampilkan di API adalah entity rmq,
	*/
	CreateQueueHumanDetection(rmqPublisherHumanDetection RmqPublisherHumanDetection) (RmqPublisherHumanDetection, error)
}

type repository struct {
	rabbitmq *amqp.Connection
}

func NewRepository(rabbitmq *amqp.Connection) *repository {
	return &repository{rabbitmq}
}

func (r *repository) CreateQueueHumanDetection(rmqPublisherHumanDetection RmqPublisherHumanDetection) (RmqPublisherHumanDetection, error) {

	ch, err := r.rabbitmq.Channel()
	if err != nil {
		return rmqPublisherHumanDetection, err
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
		return rmqPublisherHumanDetection, err
	}

	// yang dimarshal adalah entity rmq human detection
	inputReadytoMarshal := RmqPublisherHumanDetection{
		TidID:                              rmqPublisherHumanDetection.TidID,
		DateTime:                           rmqPublisherHumanDetection.DateTime,
		Person:                             rmqPublisherHumanDetection.Person,
		ConvertedFileCaptureHumanDetection: rmqPublisherHumanDetection.ConvertedFileCaptureHumanDetection,
	}

	// Convert the HumanDetection struct to JSON
	body, err := json.Marshal(inputReadytoMarshal)
	if err != nil {
		return rmqPublisherHumanDetection, err
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
		return rmqPublisherHumanDetection, err
	}

	return rmqPublisherHumanDetection, err
}
