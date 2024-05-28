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
		dates := make([]sheets.DateTasks, config.DaysToCollect)
		for i := 0; i < config.DaysToCollect; i++ {
			t := weeek.GetWeekDayTasks(current.Format("02.01.2006"), wp)
			dayTasks := weeek.UnmarshalDateTasks(t)
			sheetsTasks := converter.ConvertWeeekSheets(dayTasks)
			dates[i] = sheetsTasks
			current = current.AddDate(0, 0, 1)
		}
		project := sheets.Project{
			Dates: dates,
		}
		err := sheets.UpdateTasksData(wp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// write converted sheets tasks to sheets
		err = sheets.WriteTasksToSheets(project, wp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	w.WriteHeader(http.StatusOK)
}
