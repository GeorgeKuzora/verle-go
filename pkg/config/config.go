package config

const (
	spreadsheetID = "1zbh7UWV9NglhkgjHx5-7WnxALfvmRTIgHZhQBdkJhYE"
	credentials   = "service-account-key.json"
)


WeeekProjects := map[WeeekProjectTypes]int{
	Unknown:  0,
	IMF120:   2,
	Trobart:  14,
	Drip:     4,
	Capsule:  5,
	Assembly: 6,
}

Imf120Workplace := Workplace{
	Sheets: Sheets{
		SpreadsheetID: 
	}
}
