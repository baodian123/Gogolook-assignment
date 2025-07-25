package request

import "github.com/baodian123/Gogolook-assignment/internal/application/command"

type UpdateTaskRequest struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func (request *UpdateTaskRequest) Validate() error {
	if request.Status != 0 && request.Status != 1 {
		return ErrUnknownTaskStatus
	}

	return nil
}

func (request *UpdateTaskRequest) ToUpdateTaskInput(id string) *command.UpdateTaskInput {
	return &command.UpdateTaskInput{
		Id:     id,
		Name:   request.Name,
		Status: request.Status,
	}
}
