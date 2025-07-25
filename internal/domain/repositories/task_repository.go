package repositories

import "github.com/baodian123/Gogolook-assignment/internal/domain/entities"

type TaskRepository interface {
	Save(task *entities.Task) error
	Find(id string) (*entities.Task, error)
	FindAll() ([]*entities.Task, error)
	Update(task *entities.Task) error
	Delete(id string) error
}
