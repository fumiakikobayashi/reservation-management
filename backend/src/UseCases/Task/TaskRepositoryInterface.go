package UseCases

import (
	Domains "backend/src/Domains/Task"
	"backend/src/UseCases/Shared"
)

type TaskRepositoryInterface interface {
	GetTasks(sortType Shared.SortType, sortOrder Shared.SortOrder) (*Domains.TaskList, error)
}
