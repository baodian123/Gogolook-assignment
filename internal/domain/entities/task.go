package entities

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id       string
	Name     string
	Status   int
	CreateAt time.Time
	UpdateAt time.Time
}

func NewTask(name string, status int) *Task {
	return &Task{
		Id:       uuid.New().String(),
		Name:     name,
		Status:   status,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
}
