package taskRepo

import (
	"time"

	"github.com/eDyrr/TaskTrackerCLI/model/task"
)

type TaskRepo struct {
	repository []task.Task
}

func (tr *TaskRepo) Add(task task.Task) {
	tr.repository = append(tr.repository, task)
}

func (tr *TaskRepo) List(filter string) (tasks []task.Task) {
	if filter == "" {
		return tr.repository
	}

	for _, task := range tr.repository {
		if task.Status == filter {
			tasks = append(tasks, task)
		}
	}

	return tasks
}

func (tr *TaskRepo) Delete() {

}

func (tr *TaskRepo) Update() {

}

func (tr *TaskRepo) UpdateStatus(status string, id int) bool {

	result := tr.Search(id)

	if result == -1 {
		return false
	}

	tr.repository[result].Status = status

	var updateTime time.Time
	updateTime = time.Now()

	tr.repository[result].UpdatedAt = &updateTime

	return true
}

func (tr *TaskRepo) Search(id int) int {
	for index, task := range tr.repository {
		if task.Id == id {
			return index
		}
	}
	return -1
}
