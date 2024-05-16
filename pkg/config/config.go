package config

type Workplace struct {
	sheets Sheets
	weeek Weeek
}

type Sheets struct {
	spreadsheetID string
	readRange     string
	credentials   string
}

type Weeek struct {
	baseUsl string
	_project WeeekProjectTypes
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

WeeekProjects := map[WeeekProjectTypes]int {
	Unknown: 0,
	IMF120: 2,
	Trobart: 14,
	Drip: 4,
	Capsule: 5,
	Assembly: 6,
}
