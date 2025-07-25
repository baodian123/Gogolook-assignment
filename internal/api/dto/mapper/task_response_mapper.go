package mapper

import (
	"github.com/baodian123/Gogolook-assignment/internal/api/dto/response"
	"github.com/baodian123/Gogolook-assignment/internal/application/command"
	"github.com/baodian123/Gogolook-assignment/internal/application/common"
	"github.com/baodian123/Gogolook-assignment/internal/application/query"
)

func ToTaskListResponse(query *query.TaskListQueryResult) []*common.TaskResult {
	result := make([]*common.TaskResult, 0, len(query.Result))

	for _, t := range query.Result {
		if t == nil {
			continue
		}

		result = append(result, &common.TaskResult{
			Id:       t.Id,
			Name:     t.Name,
			Status:   t.Status,
			CreateAt: t.CreateAt,
			UpdateAt: t.UpdateAt,
		})
	}

	return result
}

func ToCreateTaskResponse(output *command.CreateTaskOutput) *response.CreateTaskResponse {
	return &response.CreateTaskResponse{
		Id: output.Id,
	}
}

func ToUpdateTaskResponse(output *command.UpdateTaskOutput) *response.UpdateTaskResponse {
	return &response.UpdateTaskResponse{
		Id:       output.Id,
		Name:     output.Name,
		Status:   output.Status,
		CreateAt: output.CreateAt,
		UpdateAt: output.UpdateAt,
	}
}
