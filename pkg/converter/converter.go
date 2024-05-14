package converter

import (
	"verle_go/pkg/sheets"
	"verle_go/pkg/weeek"
)

func ConvertWeeekSheets(wdt weeek.DateTasks) sheets.DateTasks {
	sdt := sheets.DateTasks{
		Tasks: make([]sheets.Task, 0, 5),
	}

	for _, v := range wdt.Tasks {
		sheetsTask := sheets.Task{}
		sheetsTask.Id = v.Id
		sheetsTask.Title = v.Title
		sheetsTask.Desc = v.Desc
		sheetsTask.Date = sheets.Date(v.Date)
		sdt.Tasks = append(sdt.Tasks, sheetsTask)
	}
	return sdt
}
