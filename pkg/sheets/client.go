package sheets

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const (
	spreadsheetID = "1zbh7UWV9NglhkgjHx5-7WnxALfvmRTIgHZhQBdkJhYE"
	readRange     = "Sheet1!A:D"
	credentials   = "service-account-key.json"
)

var sheetsService *sheets.Service

func InitClient() {
	// Load the Google Sheets API credentials from your JSON file.
	creds, err := os.ReadFile(credentials)
	if err != nil {
		log.Fatalf("Unable to read credentials file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(creds, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Unable to create JWT config: %v", err)
	}

	client := config.Client(context.Background())
	sheetsService, err = sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Google Sheets service: %v", err)
	}

}

func ReadData(w http.ResponseWriter, r *http.Request) {
	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, readRange).Context(r.Context()).Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Data [][]interface{} `json:"data"`
	}{
		Data: resp.Values,
	}

	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func CreateData(w http.ResponseWriter, r *http.Request) {
	// Parse request body to get data to be added.
	type RequestData struct {
		Values [][]interface{} `json:"data"`
	}

	var requestData RequestData

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	values := sheets.ValueRange{Values: requestData.Values}
	_, err = sheetsService.Spreadsheets.Values.Append(spreadsheetID, readRange, &values).ValueInputOption("RAW").Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Update Data
func UpdateData(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		RowToUpdate int           `json:"row_to_update"`
		NewData     []interface{} `json:"new_data"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update data in the Google Sheets using the Google Sheets API.
	rangeToUpdate := fmt.Sprintf("Sheet1!A%d:C%d", requestData.RowToUpdate, requestData.RowToUpdate)
	values := sheets.ValueRange{Values: [][]interface{}{requestData.NewData}}

	_, err = sheetsService.Spreadsheets.Values.Update(spreadsheetID, rangeToUpdate, &values).
		ValueInputOption("RAW").
		Context(r.Context()).
		Do()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete Data
func DeleteData(w http.ResponseWriter, r *http.Request) {
	var rowsToDelete []int

	err := json.NewDecoder(r.Body).Decode(&rowsToDelete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete data from the Google Sheet using the Google Sheets API.
	var requests []*sheets.Request

	for _, row := range rowsToDelete {
		// Define a clear request for each row to be deleted.
		requests = append(requests, &sheets.Request{
			DeleteDimension: &sheets.DeleteDimensionRequest{
				Range: &sheets.DimensionRange{
					SheetId:    0, // You might need to adjust the sheet ID.
					Dimension:  "ROWS",
					StartIndex: int64(row - 1), // Google Sheets indexes start from 0.
					EndIndex:   int64(row),
				},
			},
		})
	}

	// Execute the batch update to delete rows.
	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{Requests: requests}
	_, err = sheetsService.Spreadsheets.BatchUpdate(spreadsheetID, batchUpdateRequest).Context(r.Context()).Do()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func WriteTasksToSheets(w http.ResponseWriter, tasks DateTasks) {

	type TasksData struct {
		Values [][]interface{} `json:"data"`
	}

	var tasksData TasksData

	for _, v := range tasks.Tasks {

		var task []interface{}
		task = append(task, fmt.Sprint(v.Id))
		task = append(task, v.Title)
		if v.Desc == "" {
			task = append(task, "no_description")
		} else {
			task = append(task, v.Desc)
		}
		task = append(task, string(v.Date.toString()))
		tasksData.Values = append(tasksData.Values, task)
	}

	values := sheets.ValueRange{Values: tasksData.Values}
	_, err := sheetsService.Spreadsheets.Values.Append(spreadsheetID, readRange, &values).ValueInputOption("RAW").Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
