package src

import (
	"backend/src/Infrastructures/Repositories"
	Handlers2 "backend/src/Presentations/Handlers"
	"backend/src/Shared"
	UseCases "backend/src/UseCases/Task"
	"github.com/jinzhu/gorm"
)

type Handlers struct {
	TaskHandler Handlers2.TaskHandler
}

func NewHandlers(db *gorm.DB, logger *Shared.LoggerInterface) *Handlers {
	return &Handlers{
		TaskHandler: *injectTaskHandlerDependencies(db, logger),
	}
}

func injectTaskHandlerDependencies(db *gorm.DB, logger *Shared.LoggerInterface) *Handlers2.TaskHandler {
	taskRepository := Repositories.NewTaskRepository(db, logger)
	getTasksUseCase := UseCases.NewGetTasksUseCase(taskRepository, logger)
	return Handlers2.NewTaskHandler(
		getTasksUseCase,
		logger,
	)
}
