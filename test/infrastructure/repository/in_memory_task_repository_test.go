package repository

import (
	"testing"
	"time"

	"github.com/baodian123/Gogolook-assignment/internal/domain/entities"
	"github.com/baodian123/Gogolook-assignment/internal/infrastructure/repository"
)

func TestInMemoryTaskRepository_SaveAndFind(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()

	task := entities.NewTask("Test Task", 1)

	err := repo.Save(task)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = repo.Save(task)

	if err != repository.ErrTaskAlreadyExists {
		t.Errorf("expected ErrTaskAlreadyExists, got %v", err)
	}

	found, err := repo.Find(task.Id)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if found.Id != task.Id || found.Name != task.Name || found.Status != task.Status {
		t.Errorf("found task does not match saved task")
	}

	_, err = repo.Find("non-existent-id")

	if err != repository.ErrTaskNotFound {
		t.Errorf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestInMemoryTaskRepository_SaveNilTask(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()

	err := repo.Save(nil)

	if err != repository.ErrInvalidTaskPassed {
		t.Errorf("expected ErrInvalidTaskPassed, got %v", err)
	}
}

func TestInMemoryTaskRepository_FindAll(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()

	task1 := entities.NewTask("Task 1", 1)
	task2 := entities.NewTask("Task 2", 2)

	repo.Save(task1)
	repo.Save(task2)

	tasks, err := repo.FindAll()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(tasks))
	}
}

func TestInMemoryTaskRepository_Update(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()

	task := entities.NewTask("Task to Update", 1)
	repo.Save(task)

	err := repo.Update(nil)

	if err != repository.ErrInvalidTaskPassed {
		t.Errorf("expected ErrInvalidTaskPassed, got %v", err)
	}

	oldUpdateAt := task.UpdateAt

	// ensure UpdateAt changes
	time.Sleep(10 * time.Millisecond)

	task.Name = "Updated Name"
	task.Status = 0
	err = repo.Update(task)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	found, _ := repo.Find(task.Id)

	if found.Name != "Updated Name" || found.Status != 0 {
		t.Errorf("task not updated correctly")
	}

	if !found.UpdateAt.After(oldUpdateAt) {
		t.Errorf("UpdateAt not updated")
	}
}

func TestInMemoryTaskRepository_Delete(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()

	task := entities.NewTask("Task to Delete", 1)
	repo.Save(task)

	err := repo.Delete(task.Id)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = repo.Find(task.Id)

	if err != repository.ErrTaskNotFound {
		t.Errorf("expected ErrTaskNotFound, got %v", err)
	}

	err = repo.Delete("non-existent-id")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
