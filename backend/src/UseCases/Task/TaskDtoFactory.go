package UseCases

import (
	Domains "backend/src/Domains/Task"
	Dto "backend/src/UseCases/Dto/Task"
)

func CreateTaskDto(task *Domains.Task) Dto.TaskDto {
	return Dto.NewTaskDto(
		task.GetTaskId().GetValue(),
		task.GetName(),
		task.GetDeadline().Format(Domains.DeadlineFormat),
		task.GetIsFavorite(),
		task.GetIsCompleted(),
	)
}
