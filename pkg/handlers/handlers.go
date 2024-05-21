package handlers

import (
	"net/http"
	"time"
	"verle_go/pkg/config"
	"verle_go/pkg/converter"
	"verle_go/pkg/sheets"
	"verle_go/pkg/weeek"
)

func InitClients() {
	sheets.InitClient()
}

func RegisterHandlers() {
	http.HandleFunc("/post", postDayTasks)
}

func postDayTasks(w http.ResponseWriter, r *http.Request) {
	for _, wp := range config.Workplaces {
		currentYear, currentMonth, currentDate := time.Now().Date()
		current := time.Date(currentYear, currentMonth, currentDate, 1, 1, 1, 1, time.Local)
		sheets.UpdateTasksData(w, r, wp)
		for i := 0; i <= 7; i++ {
			t := weeek.GetWeekDayTasks(current.Format("02.01.2006"), wp)
			dayTasks := weeek.UnmarshalDateTasks(t)
			// convert to sheets tasks
			sheetsTasks := converter.ConvertWeeekSheets(dayTasks)
			// write converted sheets tasks to sheets
			sheets.WriteTasksToSheets(w, sheetsTasks, wp)
			current = current.AddDate(0, 0, 1)
		}
	}
}
