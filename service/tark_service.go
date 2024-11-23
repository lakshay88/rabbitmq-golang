package service

import (
	"encoding/json"

	"github.com/lakshay88/rabbitmq-golang/models"
	"github.com/lakshay88/rabbitmq-golang/queue"
	"github.com/lakshay88/rabbitmq-golang/repository"
	amqp "github.com/rabbitmq/amqp091-go"
)

type TaskService struct {
	repo repository.TaskRepository
	ch   *amqp.Channel
}

func NewTaskService(repo repository.TaskRepository, ch *amqp.Channel) *TaskService {
	return &TaskService{
		repo: repo,
		ch:   ch,
	}
}

func (t *TaskService) CreateTask(task *models.Task) error {
	t.repo.CreateTask(task)

	taskBytes, _ := json.Marshal(task)

	queue.PublishTask(t.ch, "task-exchange", taskBytes, task.Priority)
	return nil
}
