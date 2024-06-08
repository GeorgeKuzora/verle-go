package handlers

import (
	"html/template"
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
	http.HandleFunc("/render", render)
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
		go func(pt tasks.ProjectType) {
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
			}
			err = project.Write(project.Dates)
			if err != nil {
				log.Printf("can't write to sheets for a project %v", project)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}(pt)
	}
	w.WriteHeader(http.StatusOK)
}

func render(w http.ResponseWriter, r *http.Request) {
	type Content struct {
		Title string
		Text  string
	}
	if r.Method == "GET" {
		some_template, _ := template.ParseFiles("templates/tasks.html")
		some_content := Content{
			Title: "Это заголовок",
			Text:  "Это текст",
		}
		err := some_template.Execute(w, some_content)

		if err != nil {
			log.Println("error during page rendering")
		}
	}
}
