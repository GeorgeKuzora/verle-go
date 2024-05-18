package config

type Workplace struct {
	Sheets Sheets
	Weeek  Weeek
}

type Sheets struct {
	SpreadsheetID string
	ReadRange     string
	Credentials   string
}

type Weeek struct {
	BaseUsl string
	project WeeekProjectTypes
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
