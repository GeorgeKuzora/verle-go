package sheets

import (
	"time"
)

// Project represents all open tasks for a production center.
// Task is represented as a slice with tasks on future Dates
type Project struct {
	Dates []DateTasks `json:"dates"`
}

// DateTasks represents all tasks for a given date.
type DateTasks struct {
	Tasks []Task `json:"tasks"`
}

// Task represents a weeek task card. It includes:
// Card Id, Title, Description, Date
type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"description"`
	Date  Date   `json:"date"`
}

// Date represents Task date in proper format
type Date time.Time

func (d Date) String() string {
	t := time.Time(d)
	s := t.Format("02.01.2006")
	return s
}
