package query

import "github.com/baodian123/Gogolook-assignment/internal/application/common"

type TaskQueryResult struct {
	Result *common.TaskResult
}

type TaskListQueryResult struct {
	Result []*common.TaskResult
}
