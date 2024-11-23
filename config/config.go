package config

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
		return nil, nil, err
	}

	return conn, ch, nil
}

// Declare exchanges and queues
func SetupRabbitMQ(ch *amqp.Channel) error {
	// Set up exchanges

	{
		// Declare Direct Exchange for task routing
		err := ch.ExchangeDeclare(
			"task-exchange",
			"direct",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			failOnError(err, "Failed to declare task exchange:")
			return err
		}

		// Declare Dead-Letter Exchange
		err = ch.ExchangeDeclare(
			"dead-letter-exchange",
			"direct",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			failOnError(err, "Failed to declare dead-letter exchange:")
			return err
		}
	}

	// Declare Priority Queue for task processing
	p := amqp.Table{
		"x-max-priority":            10, // Maximum priority (0-10)
		"x-dead-letter-exchange":    "dead-letter-exchange",
		"x-dead-letter-routing-key": "dead-tasks", // Routing key for dead-letter
	}
	_, err := ch.QueueDeclare(
		"task-queue", // Queue name
		true,         // Durable
		false,        // Auto-deleted
		false,        // Exclusive
		false,        // No-wait
		p,            // Arguments
	)
	if err != nil {
		failOnError(err, "Failed to declare task queue:")
		return err
	}

	// Ques Declarations
	// Declare Dead-Letter Queue
	_, err = ch.QueueDeclare(
		"dead-letter-queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to declare dead-letter queue:", err)
		return err
	}

	// BINDING STARTED
	// Bind queues to exchanges
	err = ch.QueueBind(
		"task-queue",    // Queue name
		"task.create",   // Routing key
		"task-exchange", // Exchange name
		false,           // No-wait
		nil,             // Arguments
	)
	if err != nil {
		log.Println("Failed to bind task queue to exchange:", err)
		return err
	}

	// Bind Dead-Letter Queue to Dead-Letter Exchange
	err = ch.QueueBind(
		"dead-letter-queue",    // Queue name
		"dead-tasks",           // Routing key
		"dead-letter-exchange", // Exchange name
		false,                  // No-wait
		nil,                    // Arguments
	)
	if err != nil {
		log.Println("Failed to bind dead-letter queue to exchange:", err)
		return err
	}

	log.Println("RabbitMQ exchanges, queues, and bindings setup successfully")
	return nil
}

func failOnError(err error, msg string) {
	log.Panicf("%s: %s", msg, err)
}
