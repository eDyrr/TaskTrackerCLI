package taskRepo

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/eDyrr/TaskTrackerCLI/model/task"
)

type TaskRepo struct {
	repository []task.Task
}

func (tr *TaskRepo) LoadData() {
	fileContent, _ := ioutil.ReadFile("test.json")
	json.Unmarshal(fileContent, &tr.repository)
}

func (tr *TaskRepo) SaveData() {
	data, _ := json.MarshalIndent(tr.repository, "", "")
	ioutil.WriteFile("test.json", data, 0644)
}

func (tr *TaskRepo) Add(task *task.Task) {
	tr.repository = append(tr.repository, *task)
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

func (tr *TaskRepo) Delete(id int) {
	// tasks := make([]task.Task, 0)
	index := tr.Search(id)
	// tasks = append(tasks, tr.repository[:index]...)
	// tasks = append()
	tr.repository = append(tr.repository[:index], tr.repository[index+1:]...)
}

func (tr *TaskRepo) Update(id int, description string) {
	tr.repository[tr.Search(id)].Description = description
	var now time.Time
	now = time.Now()
	tr.repository[tr.Search(id)].UpdatedAt = &now
}

func (tr *TaskRepo) UpdateStatus(status string, id int) bool {

	result := tr.Search(id)

	if result == -1 {
		return false
	}
	// fmt.Print(result)
	tr.repository[result].Status = status

	var updateTime time.Time
	updateTime = time.Now()

	tr.repository[result].UpdatedAt = &updateTime
	// fmt.Print(tr.repository[result])
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
