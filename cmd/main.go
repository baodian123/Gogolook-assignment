package main

import (
	"github.com/baodian123/Gogolook-assignment/internal/api/controller"
	"github.com/baodian123/Gogolook-assignment/internal/application/services"
	"github.com/baodian123/Gogolook-assignment/internal/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {
	r := gin.Default()

	container := dig.New()
	container.Provide(repository.NewInMemoryTaskRepository)
	container.Provide(services.NewTaskService)
	container.Provide(controller.NewTaskController)

	container.Invoke(func(ctrl *controller.TaskController) {
		ctrl.RegisterRoutes(r)
	})

	r.Run()
}
