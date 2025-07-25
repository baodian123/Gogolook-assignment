package controller

import (
	"net/http"

	"github.com/baodian123/Gogolook-assignment/internal/api/dto/mapper"
	"github.com/baodian123/Gogolook-assignment/internal/api/dto/request"
	"github.com/baodian123/Gogolook-assignment/internal/application/interfaces"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	svc interfaces.TaskService
}

func NewTaskController(svc interfaces.TaskService) *TaskController {
	return &TaskController{svc: svc}
}

func (ctrl *TaskController) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/tasks", ctrl.GetTaskList)
	engine.POST("/tasks", ctrl.CreateTask)
	engine.PUT("/tasks/:id", ctrl.UpdateTask)
	engine.DELETE("/tasks/:id", ctrl.DeleteTask)
}

func (ctrl *TaskController) GetTaskList(ctx *gin.Context) {
	taskListQueryResult, err := ctrl.svc.GetTaskList()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Fail to get task list",
		})
	}

	ctx.JSON(http.StatusOK, mapper.ToTaskListResponse(taskListQueryResult))
}

func (ctrl *TaskController) CreateTask(ctx *gin.Context) {
	var createTaskRequest request.CreateTaskRequest

	if err := ctx.BindJSON(&createTaskRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Bad request",
		})
		return
	}

	createTaskOutput, err := ctrl.svc.CreateTask(createTaskRequest.ToCreateTaskInput())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Fail to create task",
		})
		return
	}

	ctx.JSON(http.StatusCreated, mapper.ToCreateTaskResponse(createTaskOutput))
}

func (ctrl *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid or missing task id",
		})
		return
	}

	var updateTaskRequest request.UpdateTaskRequest

	if err := ctx.BindJSON(&updateTaskRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Bad request",
		})
		return
	}

	updateTaskOutput, err := ctrl.svc.UpdateTask(updateTaskRequest.ToUpdateTaskInput(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Fail to update task",
		})
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToUpdateTaskResponse(updateTaskOutput))
}

func (ctrl *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid or missing task id",
		})
		return
	}

	err := ctrl.svc.DeleteTask(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Fail to delete task",
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
