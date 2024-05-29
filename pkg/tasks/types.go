package tasks

import (
	"time"
)

// Metadata represents information about project
// that is needed to identify project properties
type Metadata struct {
	Type   ProjectType
	Period PeriodInDays
}

// ProjectType represents type of available projects
// Usually it named after concrete production center
type ProjectType int

const (
	Unknown ProjectType = iota
	IMF120
	Trobart
	Drip
	Capsule
	Assembly
)

// PeriodInDays represents period from today
// that will be fetched and proceed
type PeriodInDays int

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

// Interface for fetching tasks data from an external service
type Fetcher interface {
	Fetch() (Tasks, error)
}

// Interface for writing tasks data to an external service
type Writer interface {
	Write() error
}
