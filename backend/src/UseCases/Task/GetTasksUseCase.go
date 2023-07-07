package UseCases

import (
	Requests "backend/src/Presentations/Requests/Task"
	"backend/src/Shared"
	Dto "backend/src/UseCases/Dto/Task"
	Shared2 "backend/src/UseCases/Shared"
)

type GetTasksUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         *Shared.LoggerInterface
}

func NewGetTasksUseCase(taskRepository TaskRepositoryInterface, logger *Shared.LoggerInterface) *GetTasksUseCase {
	return &GetTasksUseCase{
		taskRepository: taskRepository,
		logger:         logger,
	}
}

func (u *GetTasksUseCase) Execute(tasksRequest Requests.GetTasksRequest) (Dto.TaskListDto, error) {
	sortType, err := Shared2.NewSortType(tasksRequest.Sort)
	if err != nil {
		return Dto.TaskListDto{}, err
	}
	sortOrder, err := Shared2.NewSortOrder(tasksRequest.Order)
	if err != nil {
		return Dto.TaskListDto{}, err
	}

	taskList, err := u.taskRepository.GetTasks(sortType, sortOrder)
	if err != nil {
		return Dto.TaskListDto{}, err
	}

	return CreateTaskDtoList(taskList), nil
}
