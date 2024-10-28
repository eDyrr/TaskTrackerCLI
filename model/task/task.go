package task

import (
	"time"
)

type Status int

const (
	Done Status = iota
	InProgress
	ToDo
	unknown
)

var StatusName = map[Status]string{
	Done:       "done",
	InProgress: "in progress",
	ToDo:       "to do",
	unknown:    "unknown",
}

func translate(s string) Status {
	for status, value := range StatusName {
		if value == s {
			return status
		}
	}
	return unknown
}

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

// func loadTasks() []Task {
// 	var tasks []Task

// 	fileContent, _ := ioutil.ReadFile("test.json")
// 	json.Unmarshal(fileContent, &tasks)
// 	return tasks
// }

// func saveTasks(tasks []Task) {
// 	data, _ := json.MarshalIndent(tasks, "", "")
// 	ioutil.WriteFile("test.json", data, 0644)
// }

// func (task *Task) Add() {
// 	tasks := loadTasks()

// 	tasks = append(tasks, *task)

// 	saveTasks(tasks)
// }

// func List(filter string) []Task {

// 	if filter == "" {
// 		return loadTasks()
// 	}

// 	var result []Task
// 	tasks := loadTasks()

// 	for _, task := range tasks {
// 		if task.Status == filter {
// 			result = append(result, task)
// 		}
// 	}

// 	return result
// }

// func search(description string) Task {

// 	tasks := loadTasks()

// 	for _, task := range tasks {
// 		if task.Description == description {
// 			return task
// 		}
// 	}
// 	return Task{}
// }

// func update(updated Task) {
// 	tasks := loadTasks()
// 	for _, task := range tasks {
// 		if updated.Id == task.Id {
// 			task = updated
// 		}
// 	}
// 	saveTasks(tasks)
// }

// func updateStatus(description string, status string) {
// 	task := search(description)
// 	tasks := loadTasks()
// 	task.Status = status

// 	for _, value := range tasks {
// 		if value.Id == task.Id {
// 			value = task
// 		}
// 	}
// 	saveTasks(tasks)
// }

// func (task *Task) UpdateStatus(status string) {
// 	tasks := loadTasks()
// 	task.Status = status

// 	for _, value := range tasks {
// 		if value.Id == task.Id {
// 			value = *task
// 		}
// 	}
// 	saveTasks(tasks)
// }
