package main

import (
	"flag"
	"fmt"
	"time"
)

type Flag int

const (
	add Flag = iota
	delete
	update
	list
	mark_in_progress
	mark_done
)

func (f Flag) String() string {
	switch f {
	case add:
		return "add"
	case delete:
		return "delete"
	case update:
		return "update"
	case list:
		return "list"
	case mark_in_progress:
		return "mark-in-progress"
	case mark_done:
		return "mark-done"
	default:
		return "unkown"
	}
}

type Status int

const (
	done Status = iota
	todo
	in_progress
)

func (s Status) String() string {
	switch s {
	case done:
		return "done"
	case todo:
		return "todo"
	case in_progress:
		return "in-progress"
	default:
		return "unkown"
	}
}

type Task struct {
	id          int
	description string
	status      Status
	createdAt   time.Time
	updatedAt   time.Time
}

type taskRepo interface {
	GetTasks() ([]Task, error)
}

func main() {
	greeting := flag.String("Hi", "eDD", "just a greeting")
	flag.Parse()

	fmt.Printf("Hello, %s\n", *greeting)
}
