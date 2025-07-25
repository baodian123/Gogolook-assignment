package services

import (
	"github.com/baodian123/Gogolook-assignment/internal/application/command"
	"github.com/baodian123/Gogolook-assignment/internal/application/common"
	"github.com/baodian123/Gogolook-assignment/internal/application/interfaces"
	"github.com/baodian123/Gogolook-assignment/internal/application/query"
	"github.com/baodian123/Gogolook-assignment/internal/domain/entities"
	"github.com/baodian123/Gogolook-assignment/internal/domain/repositories"
)

type TaskService struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(repository repositories.TaskRepository) interfaces.TaskService {
	return &TaskService{taskRepository: repository}
}

func (svc *TaskService) GetTaskList() (*query.TaskListQueryResult, error) {
	tasks, err := svc.taskRepository.FindAll()

	if err != nil {
		return nil, err
	}

	taskListQueryResult := &query.TaskListQueryResult{
		Result: make([]*common.TaskResult, 0, len(tasks)),
	}

	for _, task := range tasks {
		taskListQueryResult.Result = append(taskListQueryResult.Result, &common.TaskResult{
			Id:       task.Id,
			Name:     task.Name,
			Status:   task.Status,
			CreateAt: task.CreateAt,
			UpdateAt: task.UpdateAt,
		})
	}
	return taskListQueryResult, nil
}

func (svc *TaskService) CreateTask(input *command.CreateTaskInput) (*command.CreateTaskOutput, error) {
	task := entities.NewTask(input.Name, input.Status)
	err := svc.taskRepository.Save(task)

	if err != nil {
		return nil, err
	}

	return &command.CreateTaskOutput{Id: task.Id}, nil
}

func (svc *TaskService) UpdateTask(input *command.UpdateTaskInput) (*command.UpdateTaskOutput, error) {
	task, err := svc.taskRepository.Find(input.Id)

	if err != nil {
		return nil, err
	}

	task.Name = input.Name
	task.Status = input.Status

	if err := svc.taskRepository.Update(task); err != nil {
		return nil, err
	}

	return &command.UpdateTaskOutput{
		TaskResult: common.TaskResult{
			Id:       task.Id,
			Name:     task.Name,
			Status:   task.Status,
			CreateAt: task.CreateAt,
			UpdateAt: task.UpdateAt,
		},
	}, nil
}

func (svc *TaskService) DeleteTask(id string) error {
	if err := svc.taskRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
