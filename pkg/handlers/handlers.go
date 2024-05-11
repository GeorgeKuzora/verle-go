package handlers

import (
	"fmt"
	"net/http"
	"time"
	"verle_go/pkg/sheets"
	"verle_go/pkg/weeek"
)

func RegisterHandlers() {
	http.HandleFunc("/read", readData)
	http.HandleFunc("/create", createData)
	http.HandleFunc("/update", updateData)
	http.HandleFunc("/delete", deleteData)
	http.HandleFunc("/getweeek", getDayTasks)
}

func readData(w http.ResponseWriter, r *http.Request) {
	sheets.ReadData(w, r)
}

func createData(w http.ResponseWriter, r *http.Request) {
	sheets.CreateData(w, r)
}

func updateData(w http.ResponseWriter, r *http.Request) {
	sheets.UpdateData(w, r)
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	sheets.DeleteData(w, r)
}

func getDayTasks(w http.ResponseWriter, r *http.Request) {
	tasks := weeek.GetWeekDayTasks("13.05.2024")
	obj := weeek.UnmarshalDateTasks(tasks)
	fmt.Printf("%+v\n", obj)
	fmt.Println(time.Time(obj.Tasks[0].Date))
}

}
