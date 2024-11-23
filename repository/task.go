package repository

import "github.com/lakshay88/rabbitmq-golang/models"

type TaskRepository interface {
	CreateTask(task *models.Task)
}
