package helper

import (
	"sync"

	"github.com/baodian123/Gogolook-assignment/internal/domain/entities"
)

type TaskSyncMap struct {
	m sync.Map
}

func (tm *TaskSyncMap) Store(key string, value entities.Task) {
	tm.m.Store(key, value)
}

func (tm *TaskSyncMap) Load(key string) (*entities.Task, bool) {
	val, ok := tm.m.Load(key)

	if !ok {
		return nil, false
	}

	task := val.(entities.Task)

	return &task, true
}

func (tm *TaskSyncMap) LoadOrStore(key string, value entities.Task) (*entities.Task, bool) {
	actual, loaded := tm.m.LoadOrStore(key, value)
	task := actual.(entities.Task)

	return &task, loaded
}

func (tm *TaskSyncMap) Delete(key string) {
	tm.m.Delete(key)
}

func (tm *TaskSyncMap) Range(f func(key string, value entities.Task) bool) {
	tm.m.Range(func(k, v any) bool {
		return f(k.(string), v.(entities.Task))
	})
}
