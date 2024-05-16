package sheets

import (
	"time"
)

// DateTasks represents all tasks for a given date.
type DateTasks struct {
	Tasks []Task `json:"tasks"`
}

// Task represents a weeek task card.
type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"` 
	Desc  string `json:"description"`
	Date  Date   `json:"date"`
}

type Date time.Time

func (d Date) toString() string {
	t := time.Time(d)
	s := t.Format("02.01.2006")
	return s
}
