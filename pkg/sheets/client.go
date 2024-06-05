package sheets

import (
	"context"
	"fmt"
	"log"
	"os"
	"verle_go/pkg/tasks"

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

var projects = map[tasks.ProjectType]project{
	tasks.IMF120: project{
		SpreadsheetID: spreadsheetID,
		Range:         "IMF!A:D",
		UpdateRange:   "IMF!A1:D200",
		SheetID:       1330258137,
	},
	tasks.Trobart: project{
		SpreadsheetID: spreadsheetID,
		Range:         "TROBART!A:D",
		UpdateRange:   "TROBART!A1:D200",
		SheetID:       1738797376,
	},
	tasks.Drip: project{
		SpreadsheetID: spreadsheetID,
		Range:         "DRIPS!A:D",
		UpdateRange:   "DRIPS!A1:D200",
		SheetID:       612152640,
	},
	tasks.Capsule: project{
		SpreadsheetID: spreadsheetID,
		Range:         "CAPSULES!A:D",
		UpdateRange:   "CAPSULES!A1:D200",
		SheetID:       1199560039,
	},
	tasks.Assembly: project{
		SpreadsheetID: spreadsheetID,
		Range:         "ASSEMBLY!A:D",
		UpdateRange:   "ASSEMBLY!A1:D200",
		SheetID:       1355663488,
	},
}

type project struct {
	SpreadsheetID string
	Range         string
	UpdateRange   string
	SheetID       int
}

type TaskWriter struct {
	Project tasks.ProjectType
}

func (tw TaskWriter) Write(dates []tasks.Tasks) error {
	err := tw.cleanRange()
	if err != nil {
		log.Printf("can't clean spreadsheet for project %v", tw.Project)
		return err
	}

	type TasksData struct {
		Values [][]interface{} `json:"data"`
	}

	var tasksData TasksData

	for _, d := range dates {
		for _, v := range d.Tasks {

			var task []interface{}
			task = append(task, fmt.Sprint(v.Id))
			task = append(task, v.Title)
			if v.Desc == "" {
				task = append(task, "no_description")
			} else {
				task = append(task, v.Desc)
			}
			ds, err := v.Date.String()
			if err != nil {
				log.Printf("can't covert date to string for %v", v.Date)
				task = append(task, "")
			} else {
				task = append(task, ds)
			}
			tasksData.Values = append(tasksData.Values, task)
		}
	}

	values := sheets.ValueRange{Values: tasksData.Values}
	p, ok := projects[tw.Project]
	if ok == false {
		log.Printf("can't find sheets project from a project %v", tw.Project)
		return fmt.Errorf("can't find sheets project from a project %v", tw.Project)
	}
	_, err = sheetsService.Spreadsheets.Values.Append(p.SpreadsheetID, p.Range, &values).ValueInputOption("RAW").Do()
	if err != nil {
		log.Printf("error during writing data in sheet with id %s, list %s", p.SpreadsheetID, p.Range)
		return fmt.Errorf("error during writing data in sheet with id %s, list %s", p.SpreadsheetID, p.Range)
	}
	return nil
}

func (tw TaskWriter) cleanRange() error {
	var updateData [][]interface{}
	var emptyData []interface{}

	for i := 0; i < 4; i++ {
		emptyData = append(emptyData, "")
	}
	for i := 0; i < 200; i++ {
		updateData = append(updateData, emptyData)
	}
	// Update data in the Google Sheets using the Google Sheets API.
	values := sheets.ValueRange{Values: updateData}

	p, ok := projects[tw.Project]
	if ok == false {
		log.Printf("can't find sheets project from a project %v", tw.Project)
		return fmt.Errorf("can't find sheets project from a project %v", tw.Project)
	}

	_, err := sheetsService.Spreadsheets.Values.Update(p.SpreadsheetID, p.UpdateRange, &values).
		ValueInputOption("RAW").
		Context(context.Background()).
		Do()
	if err != nil {
		log.Printf("error during update sheet with id %s, list %s", p.SpreadsheetID, p.Range)
		return fmt.Errorf("can't update sheet with id %s", p.SpreadsheetID)
	}
	return nil
}
