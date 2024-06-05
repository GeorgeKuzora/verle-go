package handlers

import (
	"log"
	"net/http"
	"verle_go/pkg/sheets"
	"verle_go/pkg/tasks"
	"verle_go/pkg/weeek"
)

func InitClients() {
	sheets.InitClient()
}

func RegisterHandlers() {
	http.HandleFunc("/tasks", writeTasksToSheets)
}

func writeTasksToSheets(w http.ResponseWriter, r *http.Request) {
	periodInDays := tasks.PeriodInDays(10)
	dates, err := periodInDays.GetDatesFromToday()
	if err != nil {
		log.Printf("can't gets dates to process")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	projectTypes := []tasks.ProjectType{
		tasks.IMF120,
		tasks.Trobart,
		tasks.Drip,
		tasks.Capsule,
		tasks.Assembly,
	}

	for _, pt := range projectTypes {
		fetcher := weeek.TaskFetcher{
			Project: pt,
		}
		writer := sheets.TaskWriter{
			Project: pt,
		}
		project := tasks.Project{
			TasksFetcher: fetcher,
			TasksWriter:  writer,
		}
		err := project.Fetch(dates)
		if err != nil {
			log.Printf("can't fetch from weeek for a project %v", project)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			break
		}
		err = project.Write(project.Dates)
		if err != nil {
			log.Printf("can't write to sheets for a project %v", project)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			break
		}
	}
	w.WriteHeader(http.StatusOK)
}
