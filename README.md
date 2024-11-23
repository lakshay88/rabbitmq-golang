# RabbitMQ Task Management System with Golang

Golang with RabbitMQ as the message broker. It demonstrates how to produce, consume, and handle messages with RabbitMQ, including task prioritization and dead-letter queue handling for failed tasks.

# Featurn 
 1. Task Producer: Adds tasks to the RabbitMQ task queue.
 2. Task Consumer: Processes tasks from the queue.
 3. Dead-Letter Queue: Automatically routes failed tasks to a separate queue for further analysis or retries.
 4. In-Memory Task Repository: Manages tasks without requiring a database, suitable for testing and learning purposes.
 5. Priority Queues: Supports task prioritization with RabbitMQ.
 6. Random Failure Simulation: Simulates task failures for testing dead-letter functionality.


# RabbitMQ Setup with Docker

`docker run -d --hostname rabbitmq --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management`
RabbitMQ Management Console: http://localhost:15672/
Default credentials:
    Username: guest
    Password: guest

# Installation 

1. Clone
`git clone https://github.com/lakshay88/rabbitmq-golang.git`
`cd rabbitmq-golang`

2. Install dependencies:
`go mod tidy`

3. Run app
`go run main.go`

