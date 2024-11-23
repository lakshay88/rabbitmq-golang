package queue

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishTask(ch *amqp.Channel, exchange string, taskData []byte, priority int) error {
	// Publish task message with routing key and priority

	err := ch.Publish(
		exchange,
		"task.create",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        taskData,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish task: %s", err)
	}

	return nil
}
