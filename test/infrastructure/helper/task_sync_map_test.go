package helper

import (
	"testing"

	"github.com/baodian123/Gogolook-assignment/internal/domain/entities"
	"github.com/baodian123/Gogolook-assignment/internal/infrastructure/helper"
)

func TestTaskSyncMap_StoreAndLoad(t *testing.T) {
	tsm := &helper.TaskSyncMap{}
	task := entities.Task{Id: "1", Name: "Test Task", Status: 0}
	tsm.Store(task.Id, task)

	loaded, ok := tsm.Load("1")

	if !ok {
		t.Fatalf("expected to load task with id 1")
	}

	if loaded.Name != "Test Task" {
		t.Errorf("expected name 'Test Task', got %s", loaded.Name)
	}

	if loaded.Status != 0 {
		t.Errorf("expected status '0', got %d", loaded.Status)
	}
}

func TestTaskSyncMap_LoadNotFound(t *testing.T) {
	tsm := &helper.TaskSyncMap{}

	_, ok := tsm.Load("not-exist")

	if ok {
		t.Errorf("expected not to find task with id 'not-exist'")
	}
}

func TestTaskSyncMap_Delete(t *testing.T) {
	tsm := &helper.TaskSyncMap{}

	task := entities.Task{Id: "2", Name: "To Delete", Status: 1}
	tsm.Store(task.Id, task)
	tsm.Delete("2")

	val, ok := tsm.Load("2")

	if ok {
		t.Errorf("expected task with id '2' to be deleted")
	}

	if val != nil {
		t.Errorf("expected task with id '2' to be deleted and cannot be found")
	}
}

func TestTaskSyncMap_Range(t *testing.T) {
	tsm := &helper.TaskSyncMap{}

	task1 := entities.Task{Id: "a", Name: "A", Status: 0}
	task2 := entities.Task{Id: "b", Name: "B", Status: 1}
	tsm.Store(task1.Id, task1)
	tsm.Store(task2.Id, task2)

	found := map[string]bool{}
	tsm.Range(func(key string, value entities.Task) bool {
		found[key] = true
		return true
	})

	if !found["a"] || !found["b"] {
		t.Errorf("expected to find tasks in range, got: %v", found)
	}
}
