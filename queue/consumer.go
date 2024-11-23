package queue

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Consume tasks from the queue
func ConsumeTasks(ch *amqp.Channel, queueName string) error {
	log.Printf("Listening to queue: %s", queueName)
	mssg, err := ch.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err)
	}
	log.Println("Consumer registered successfully!")
	fmt.Println("================================================")
	fmt.Println("here")
	fmt.Println("================================================")

	go func() {
		for msg := range mssg {
			log.Printf("Received a message: %s", msg.Body)

			if rand.Intn(2) == 0 {
				log.Printf("Simulating failure for task: %s", msg.Body)
				msg.Nack(false, false)
				continue
			}

			err := processTask(msg.Body)
			if err != nil {
				log.Printf("Task processing failed: %v", err)
				msg.Nack(false, false) // Move to dead-letter queue if configured
				continue
			}
			msg.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for messages from queue: %s. To exit press CTRL+C", queueName)
	return nil
}

// Process dead-letter messages
func ProcessDeadLetterMessages(ch *amqp.Channel, deadLetterQueue string) error {
	// Handle messages in dead-letter queue
	msgs, err := ch.Consume(
		deadLetterQueue, // Dead-letter queue name
		"",              // Consumer tag
		false,           // Auto-acknowledge
		false,           // Exclusive
		false,           // No-local
		false,           // No-wait
		nil,             // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer for dead-letter queue: %v", err)
	}
	go func() {
		for msg := range msgs {
			log.Printf("Dead-letter message: %s", msg.Body)

			// Process the message (e.g., retry logic, alert, etc.)
			err := handleDeadLetter(msg.Body)
			if err != nil {
				log.Printf("Failed to handle dead-letter message: %v", err)
				msg.Nack(false, false)
				continue
			}
			msg.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for dead-letter messages from queue: %s. To exit press CTRL+C", deadLetterQueue)
	return nil
}

func processTask(body []byte) error {
	log.Printf("Processing task: %s", string(body))
	time.Sleep(500 * time.Millisecond)
	if string(body) == "fail" {
		return fmt.Errorf("simulated task failure")
	}
	return nil
}

func handleDeadLetter(body []byte) error {
	log.Printf("Handling dead-letter message: %s", string(body))
	return nil
}
