package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eDyrr/TaskTrackerCLI/model/task"
	taskRepo "github.com/eDyrr/TaskTrackerCLI/model/taskRepo"
)

var repo taskRepo.TaskRepo
var settings Settings

type arrayFlags []string

func (a *arrayFlags) String() string {
	return strings.Join(*a, ",")
}

func (a *arrayFlags) Set(value string) error {
	*a = append(*a, value)
	return nil
}

type Flag int

const (
	add Flag = iota
	delete
	update
	list
	list_done
	list_in_progress
	list_to_do
	mark_in_progress
	mark_done
	unknown
)

var FlagName = map[Flag]string{
	add:              "add",
	delete:           "delete",
	update:           "update",
	list:             "list",
	list_done:        "list-done",
	list_in_progress: "list-in-progress",
	list_to_do:       "list-to-do",
	mark_in_progress: "mark-in-progress",
	mark_done:        "mark-done",
	unknown:          "unknown",
}

func translate(s string) Flag {
	for flag, value := range FlagName {
		if value == s {
			return flag
		}
	}
	return unknown
}

type Filters []string

func (f *Filters) set(value string) error {
	*f = append(*f, value)
	return nil
}

var filter Filters

// func (f Flag) String() string {
// 	switch f {
// 	case add:
// 		return "add"
// 	case delete:
// 		return "delete"
// 	case update:
// 		return "update"
// 	case list:
// 		return "list"
// 	case mark_in_progress:
// 		return "mark-in-progress"
// 	case mark_done:
// 		return "mark-done"
// 	default:
// 		return "unkown"
// 	}
// }

type Settings struct {
	ID int `json:"id_count"`
}

func incrementIdCount() int {

	settings = loadSettings("settings.json")
	settings.ID++

	return settings.ID
}

func loadSettings(fileName string) Settings {
	var settings Settings

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return settings
	}

	fileContent, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("error reading settings file:", err)
		return settings
	}

	err = json.Unmarshal(fileContent, &settings)
	if err != nil {
		fmt.Println("err parsing JSON:", err)
	}
	return settings
}

func saveSettings(filename string, settings Settings) {
	jsonData, err := json.MarshalIndent(settings, "", "")

	if err != nil {
		fmt.Println("error converting settings to JSON:", err)
		return
	}

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("error writing settings file:", err)
	}
}

func coordinator() {
	Flag := flag.String("task", "", "command")
	flag.Parse()
	content := flag.Arg(0)
	// fmt.Println(*Flag)
	// fmt.Println(content)

	t := new(task.Task)

	input := translate(*Flag)
	// fmt.Print(input)
	id64, _ := strconv.ParseInt(content, 10, 0)
	id := int(id64)
	// id = 1
	// fmt.Printf("id %d", id)
	switch input {
	case add:
		t = &task.Task{
			Id:          incrementIdCount(),
			Description: content,
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		}

		repo.Add(t)
	case delete:
		repo.Delete(id)
	case update:
		repo.Update(id, content)
	case list:
		fmt.Print(repo.List(content))
	case list_done:
		repo.List(content)
	case list_in_progress:
		repo.List(content)
	case list_to_do:
		repo.List(content)
	case mark_in_progress:
		repo.UpdateStatus("in progress", id)
	case mark_done:
		repo.UpdateStatus("done", id)
	case unknown:
	}
}

func startup() {
	settings = loadSettings("settings.json")
	repo.LoadData()
}

func terminate() {
	saveSettings("settings.json", settings)
	repo.SaveData()
}

func main() {

	startup()
	coordinator()
	terminate()

	// var tags arrayFlags

	// flag.Var(&tags, "tag", "provide n tags")

	// flag.Parse()

	// fmt.Println("Tags:", tags)

	// t := task.Task{
	// 	Id:          2,
	// 	Description: "some description",
	// 	Status:      task.StatusName[task.Done],
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   (*time.Time)(nil),
	// }

	// for _, task := range tasks {
	// 	fmt.Println(task)
	// }
}
