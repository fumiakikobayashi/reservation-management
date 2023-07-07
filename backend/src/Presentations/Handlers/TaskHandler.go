package Handlers

import (
	Requests2 "backend/src/Presentations/Requests/Task"
	"backend/src/Shared"
	UseCases "backend/src/UseCases/Task"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskHandler struct {
	getTasksUseCase UseCases.GetTasksUseCase
	logger          Shared.LoggerInterface
}

func NewTaskHandler(
	getTasksUseCase *UseCases.GetTasksUseCase,
	logger *Shared.LoggerInterface,
) *TaskHandler {
	return &TaskHandler{
		getTasksUseCase: *getTasksUseCase,
		logger:          *logger,
	}
}

func (c *TaskHandler) GetTasks(ctx echo.Context) error {
	var tasksRequest Requests2.GetTasksRequest
	if err := ctx.Bind(&tasksRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	taskListDto, err := c.getTasksUseCase.Execute(tasksRequest)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, taskListDto)
}
