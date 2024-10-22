package task

import (
	"fmt"
	"time"
)

type Task struct {
	id          int
	description string
	status      string
	createdAt   time.Time
	updatedAt   time.Time
}

func Add(s *string) {
	fmt.Print(s)
}

func update(s *string) {

}

func delete(s *string) {

}
