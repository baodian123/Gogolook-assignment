package common

import "time"

type TaskResult struct {
	Id       string
	Name     string
	Status   int
	CreateAt time.Time
	UpdateAt time.Time
}
