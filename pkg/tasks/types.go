package tasks

import (
	"time"
)

type Fetcher interface {
	Fetch() (Tasks, error)
}

type Writer interface {
	Write() error
}

type WeeekProjectTypes int

const (
	Unknown WeeekProjectTypes = iota
	IMF120
	Trobart
	Drip
	Capsule
	Assembly
)

type Metadata struct {
}

// Project represents all open tasks for a production center.
// Task is represented as a slice with tasks on future Dates
type Project struct {
	Dates        []Tasks `json:"dates"`
	TasksFetcher Fetcher
	TasksWriter  Writer
}

// Tasks represents all tasks for a given date.
type Tasks struct {
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
