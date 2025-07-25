package service

import (
	"testing"

	"github.com/baodian123/Gogolook-assignment/internal/application/command"
	"github.com/baodian123/Gogolook-assignment/internal/application/services"
	"github.com/baodian123/Gogolook-assignment/internal/infrastructure/repository"
)

func setupTaskService() *services.TaskService {
	repo := repository.NewInMemoryTaskRepository()
	return services.NewTaskService(repo).(*services.TaskService)
}

func TestCreateTask(t *testing.T) {
	ts := setupTaskService()

	input := &command.CreateTaskInput{Name: "Test Task", Status: 1}
	output, err := ts.CreateTask(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if output.Id == "" {
		t.Errorf("expected non-empty task ID")
	}
}

func TestGetTaskList(t *testing.T) {
	ts := setupTaskService()

	_, _ = ts.CreateTask(&command.CreateTaskInput{Name: "Task 1", Status: 0})
	_, _ = ts.CreateTask(&command.CreateTaskInput{Name: "Task 2", Status: 1})

	result, err := ts.GetTaskList()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result.Result) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(result.Result))
	}
}

func TestUpdateTask(t *testing.T) {
	ts := setupTaskService()

	createOut, _ := ts.CreateTask(&command.CreateTaskInput{Name: "To Update", Status: 1})

	updateInput := &command.UpdateTaskInput{
		Id:     createOut.Id,
		Name:   "Updated Name",
		Status: 0,
	}

	updateOut, err := ts.UpdateTask(updateInput)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if updateOut.Name != "Updated Name" || updateOut.Status != 0 {
		t.Errorf("task not updated correctly")
	}
}

func TestDeleteTask(t *testing.T) {
	ts := setupTaskService()

	createOut, _ := ts.CreateTask(&command.CreateTaskInput{Name: "To Delete", Status: 1})

	err := ts.DeleteTask(createOut.Id)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = ts.DeleteTask(createOut.Id)

	if err != nil {
		t.Errorf("expected no error on deleting non-existent task, got %v", err)
	}
}
