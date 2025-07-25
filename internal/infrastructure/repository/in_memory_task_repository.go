package repository

import (
	"time"

	"github.com/baodian123/Gogolook-assignment/internal/domain/entities"
	"github.com/baodian123/Gogolook-assignment/internal/domain/repositories"
	"github.com/baodian123/Gogolook-assignment/internal/infrastructure/helper"
)

// taskMapper use id as key, task object as value
// implemented by sync.map to ensure thread-safe
type InMemoryTaskRepository struct {
	taskMapper helper.TaskSyncMap
}

func NewInMemoryTaskRepository() repositories.TaskRepository {
	return &InMemoryTaskRepository{}
}

func (repo *InMemoryTaskRepository) Save(task *entities.Task) error {
	if task == nil {
		return ErrInvalidTaskPassed
	}

	_, loaded := repo.taskMapper.LoadOrStore(task.Id, *task)

	if loaded {
		return ErrTaskAlreadyExists
	}

	return nil
}

func (repo *InMemoryTaskRepository) Find(id string) (*entities.Task, error) {
	task, loaded := repo.taskMapper.Load(id)

	if !loaded {
		return nil, ErrTaskNotFound
	}

	return task, nil
}

func (repo *InMemoryTaskRepository) FindAll() ([]*entities.Task, error) {
	var tasks []*entities.Task

	repo.taskMapper.Range(func(key string, value entities.Task) bool {
		task := value
		tasks = append(tasks, &task)
		return true
	})

	return tasks, nil
}

func (repo *InMemoryTaskRepository) Update(task *entities.Task) error {
	if task == nil {
		return ErrInvalidTaskPassed
	}

	task.UpdateAt = time.Now()
	repo.taskMapper.Store(task.Id, *task)

	return nil
}

func (repo *InMemoryTaskRepository) Delete(id string) error {
	// delete an key which not exist in map is safe
	// if require warning log, add load before delete to check if key is exist or not
	repo.taskMapper.Delete(id)

	return nil
}
