package Repositories

import (
	Domains2 "backend/src/Domains/Task"
	"backend/src/Infrastructures/Models"
	Shared2 "backend/src/Shared"
	Shared3 "backend/src/UseCases/Shared"
	UseCase "backend/src/UseCases/Task"
	"fmt"
	"github.com/jinzhu/gorm"
)

type taskRepository struct {
	db     *gorm.DB
	logger *Shared2.LoggerInterface
}

func NewTaskRepository(db *gorm.DB, logger *Shared2.LoggerInterface) UseCase.TaskRepositoryInterface {
	return &taskRepository{
		db:     db,
		logger: logger,
	}
}

func (r *taskRepository) GetTasks(sortType Shared3.SortType, sortOrder Shared3.SortOrder) (*Domains2.TaskList, error) {
	var taskModels []Models.TaskModel
	var sortColumn string
	taskList := Domains2.NewTaskList()

	switch sortType {
	case Shared3.Name:
		sortColumn = "name"
	case Shared3.Deadline:
		sortColumn = "deadline"
	case Shared3.Favorite:
		sortColumn = "is_favorite"
	default:
		sortColumn = "id"
	}

	if err := r.db.Table("tasks").Order(fmt.Sprintf("%s %s", sortColumn, sortOrder.GetValue())).Find(&taskModels).Error; err != nil {
		return taskList, err
	}

	for _, taskModel := range taskModels {
		task, _ := Domains2.CreateTask(taskModel)
		if err := taskList.Push(task); err != nil {
			return taskList, Shared2.NewSampleError("001-001", "doSomethingでエラー発生")
		}
	}

	return taskList, nil
}
