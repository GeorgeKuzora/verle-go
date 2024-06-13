package tasks

import (
	"encoding/json"
	"errors"
	"log"
	"time"
)

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

func (pid *PeriodInDays) GetDatesFromToday() ([]Date, error) {
	dates := make([]Date, *pid)
	current := time.Now()
	for i := 0; i < int(*pid); i++ {
		dates[i] = Date(current)
		current = current.AddDate(0, 0, 1)
	}
	return dates, nil
}

// Project represents all open tasks for a production center.
// Task is represented as a slice with tasks on future Dates
type Project struct {
	Dates        []Tasks `json:"dates"`
	TasksFetcher Fetcher
	TasksWriter  Writer
}

func (p *Project) Fetch(dates []Date) error {
	if p == nil {
		log.Println("expected Project but received nil")
		return errors.New("expected Project but received nil")
	}
	tasksForDates, err := p.TasksFetcher.Fetch(dates)
	if err != nil {
		log.Printf("can't fetch tasks for Project: %v, from Fetcher: %v", p, p.TasksFetcher)
		return err
	}
	p.Dates = tasksForDates
	err = p.fetchSubTasks()
	if err != nil {
		log.Printf("can't fetch subtasks for Project: %v, from Fetcher: %v", p, p.TasksFetcher)
	}
	return nil
}

func (p *Project) fetchSubTasks() error {
	if p == nil {
		log.Println("expected Project but received nil")
		return errors.New("expected Project but received nil")
	}
	count := 0
	subTasksPerDate := make(map[int]Tasks)
	for _, date := range p.Dates {
		for _, task := range date.Tasks {
			subTasks, err := task.fetchSubTasks(p.TasksFetcher)
			if err != nil {
				log.Printf("can't fetch subTasks for a task %v", task)
				continue
			}
			if len(subTasks) > 0 {
				subTasksPerDate[count] = Tasks{Tasks: subTasks}
				count += 1
			}
		}
	}

	for _, v := range subTasksPerDate {
		p.Dates = append(p.Dates, v)
	}
	return nil
}

func (p *Project) Write(dates []Tasks) error {
	if p == nil {
		log.Println("expected Project but received nil")
		return errors.New("expected Project but received nil")
	}
	err := p.TasksWriter.Write(dates)
	if err != nil {
		log.Printf("can't write tasks for Project: %v, to Writer: %v", p, p.TasksWriter)
		return err
	}
	return nil
}

// Tasks represents all tasks for a given date.
type Tasks struct {
	Tasks []Task `json:"tasks"`
}

// Task represents a weeek task card. It includes:
// Card Id, Title, Description, Date
type Task struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Desc       string `json:"description"`
	Date       Date   `json:"date"`
	SubTaskIDs []int  `json:"subTasks"`
}

func (t *Task) fetchSubTasks(f Fetcher) ([]Task, error) {
	if t == nil {
		log.Println("expected Task but received nil")
		return nil, errors.New("expected Task but received nil")
	}
	subTasks := make([]Task, len(t.SubTaskIDs))
	for i, taskId := range t.SubTaskIDs {
		subTask, err := f.FetchById(taskId)
		if err != nil {
			log.Printf("can't fetch a subtask with id %d from a task with id %d", taskId, t.Id)
			continue
		}
		subTasks[i] = subTask
	}
	return subTasks, nil
}

// Date represents Task date in proper format
type Date time.Time

const dateFormat string = "02.01.2006"

func (d *Date) String() (string, error) {
	if d == nil {
		log.Println("expected Date but received nil")
		return "", errors.New("expected Date but received nil")
	}
	t := time.Time(*d)
	return t.Format(dateFormat), nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := t.Format(dateFormat)
	return json.Marshal(formatted)
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	t, err := time.Parse(dateFormat, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// Interface for fetching tasks data from an external service
type Fetcher interface {
	Fetch(dates []Date) ([]Tasks, error)
	FetchById(id int) (Task, error)
}

// Interface for writing tasks data to an external service
type Writer interface {
	Write(dates []Tasks) error
}
