package config

type Workplace struct {
	SheetsTable  Sheets
	WeeekProject Weeek
}

type Sheets struct {
	SpreadsheetID string
	Range         string
	Credentials   string
}

type Weeek struct {
	Project       WeeekProjectTypes
	ProjectNumber int
}

type WeeekProjectTypes int

const (
	Unknown WeeekProjectTypes = iota
	IMF120
	Trobart
	Drip
	Capsule
	Assembly
)
