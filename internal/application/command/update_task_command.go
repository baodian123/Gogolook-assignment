package command

import "github.com/baodian123/Gogolook-assignment/internal/application/common"

type UpdateTaskInput struct {
	Id     string
	Name   string
	Status int
}

type UpdateTaskOutput struct {
	common.TaskResult
}
