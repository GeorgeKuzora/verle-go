package handlers

import (
	"fmt"
	"net/http"
	"time"
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
	t := weeek.GetWeekDayTasks("16.05.2024")
	dayTasks := weeek.UnmarshalDateTasks(t)
	// convert to sheets tasks
	sheetsTasks := converter.ConvertWeeekSheets(dayTasks)
	// write converted sheets tasks to sheets
	sheets.WriteTasksToSheets(w, sheetsTasks)
}
