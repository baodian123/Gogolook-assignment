package request

import "github.com/baodian123/Gogolook-assignment/internal/application/command"

type CreateTaskRequest struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func (request *CreateTaskRequest) Validate() error {
	if request.Status != 0 && request.Status != 1 {
		return ErrUnknownTaskStatus
	}

	return nil
}

func (request *CreateTaskRequest) ToCreateTaskInput() *command.CreateTaskInput {
	return &command.CreateTaskInput{
		Name:   request.Name,
		Status: request.Status,
	}
}
