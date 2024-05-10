package weeek

// import "time"

// Task represents a weeek task card.
type Task struct {
	Id    int       `json:"id"`
	Title string    `json: "title"`
	Desc  string    `json:"description"`
	// Date  time.Time `json:"date"`
}

// DateTasks represents all tasks for a given date.
type DateTasks struct {
	Tasks []Task `json:"tasks"`
}
