package handlers

import (
	"net/http"
	"verle_go/pkg/sheets"
)

func RegisterHandlers() {
	http.HandleFunc("/read", readData)
	http.HandleFunc("/create", createData)
	http.HandleFunc("/update", updateData)
	http.HandleFunc("/delete", deleteData)
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
