package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lakshay88/rabbitmq-golang/config"
	"github.com/lakshay88/rabbitmq-golang/models"
	"github.com/lakshay88/rabbitmq-golang/queue"
	"github.com/lakshay88/rabbitmq-golang/repository"
	"github.com/lakshay88/rabbitmq-golang/service"
)

func main() {
	conn, ch, _ := config.ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	// Setup RabbitMQ (queues, exchanges)
	config.SetupRabbitMQ(ch)

	go queue.ConsumeTasks(ch, "task-queue")

	// TODO trt
	// Uncomment Below code to start processing dead-letter queuey
	// go queue.ProcessDeadLetterMessages(ch, "dead-letter-queue")

	taskRepo := repository.NewInMemoryTaskRepository()
	taskService := service.NewTaskService(taskRepo, ch)

	addRandomTasks(taskService, 100)

	forever := make(chan bool)
	<-forever
}

func addRandomTasks(taskService *service.TaskService, count int) {
	rand.Seed(int64(time.Now().Second()))

	for i := 0; i < count; i++ {
		task := &models.Task{
			Title:       fmt.Sprintf("Task #%d", i+1),
			Description: fmt.Sprintf("Description for task #%d", i+1),
			Priority:    rand.Intn(10) + 1,
			Status:      "pending",
		}

		err := taskService.CreateTask(task)
		if err != nil {
			fmt.Printf("Error creating task #%d: %v\n", i+1, err)
		} else {
			fmt.Printf("Task #%d created and enqueued successfully!\n", i+1)
		}
	}
}
