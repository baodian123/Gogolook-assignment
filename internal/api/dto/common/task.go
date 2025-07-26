package common

import "time"

type TaskResponse struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Status   int       `json:"status"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}
