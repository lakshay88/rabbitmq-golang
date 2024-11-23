package repository

import (
	"sync"

	"github.com/lakshay88/rabbitmq-golang/models"
)

type InMemoryTaskRepo struct {
	data map[int64]*models.Task
	mu   sync.Mutex
	id   int64
}

func NewInMemoryTaskRepository() *InMemoryTaskRepo {
	return &InMemoryTaskRepo{
		data: make(map[int64]*models.Task),
	}
}
func (m *InMemoryTaskRepo) CreateTask(task *models.Task) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.id++
	task.ID = m.id
	m.data[task.ID] = task
}
