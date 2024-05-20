package handlers

import (
	"net/http"
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
		t := weeek.GetWeekDayTasks("16.05.2024", wp)
		dayTasks := weeek.UnmarshalDateTasks(t)
		// convert to sheets tasks
		sheetsTasks := converter.ConvertWeeekSheets(dayTasks)
		// write converted sheets tasks to sheets
		sheets.WriteTasksToSheets(w, sheetsTasks, wp)
	}
}
