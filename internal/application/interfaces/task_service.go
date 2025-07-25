package interfaces

import (
	"github.com/baodian123/Gogolook-assignment/internal/application/command"
	"github.com/baodian123/Gogolook-assignment/internal/application/query"
)

type TaskService interface {
	GetTaskList() (*query.TaskListQueryResult, error)
	CreateTask(input *command.CreateTaskInput) (*command.CreateTaskOutput, error)
	UpdateTask(input *command.UpdateTaskInput) (*command.UpdateTaskOutput, error)
	DeleteTask(id string) error
}
