package response

import "time"

type UpdateTaskResponse struct {
	Id       string
	Name     string
	Status   int
	CreateAt time.Time
	UpdateAt time.Time
}
