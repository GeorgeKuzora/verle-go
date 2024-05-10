package weeek

import (
	"encoding/json"
	"time"
)

// Task represents a weeek task card.
type Task struct {
	Id    int       `json:"id"`
	Title string    `json: "title"`
	Desc  string    `json:"description"`
	Date  WeeekDate `json:"date"`
}

// DateTasks represents all tasks for a given date.
type DateTasks struct {
	Tasks []Task `json:"tasks"`
}

type WeeekDate time.Time

const weeek_date_format string = "02.01.2006"

func (d WeeekDate) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := t.Format(weeek_date_format)
	return json.Marshal(formatted)
}

func (d *WeeekDate) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	t, err := time.Parse(weeek_date_format, s)
	if err != nil {
		return err
	}
	*d = WeeekDate(t)
	return nil
}
