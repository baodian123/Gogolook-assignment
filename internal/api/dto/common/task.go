package common

import "time"

type TaskResponse struct {
	Id       string
	Name     string
	Status   int
	CreateAt time.Time
	UpdateAt time.Time
}
