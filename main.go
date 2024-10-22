package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/eDyrr/TaskTrackerCLI/model/task"
)

type Flag int

const (
	add Flag = iota
	delete
	update
	list
	mark_in_progress
	mark_done
	unknown
)

var FlagName = map[Flag]string{
	add:              "add",
	delete:           "delete",
	update:           "update",
	list:             "list",
	mark_in_progress: "mark in progress",
	mark_done:        "mark done",
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
	ID int `json:id_count`
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

func main() {

	settingsFile := "settings.json"

	settings := loadSettings(settingsFile)

	settings.ID = 2

	saveSettings(settingsFile, settings)
	// Flag := flag.String("task", "", "command")
	// flag.Parse()
	// content := flag.Arg(0)
	// fmt.Println(*Flag)
	// fmt.Println(content)

	t := task.Task{
		Id:          2,
		Description: "some description",
		Status:      task.StatusName[task.Done],
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Add()
	t.UpdateStatus("done")

	tasks := task.List()

	for _, task := range tasks {
		fmt.Println(task)
	}
}
