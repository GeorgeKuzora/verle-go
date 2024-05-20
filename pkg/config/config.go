package config

const (
	spreadsheetID = "1zbh7UWV9NglhkgjHx5-7WnxALfvmRTIgHZhQBdkJhYE"
	credentials   = "service-account-key.json"
)

var WeeekProjects = map[WeeekProjectTypes]int{
	Unknown:  0,
	IMF120:   2,
	Trobart:  14,
	Drip:     4,
	Capsule:  5,
	Assembly: 6,
}

var Imf120Workplace = Workplace{
	SheetsTable: Sheets{
		SpreadsheetID: spreadsheetID,
		Range:         "IMF!A:D",
	},

	WeeekProject: Weeek{
		Project:       IMF120,
		ProjectNumber: WeeekProjects[IMF120],
	},
}

var TrobartWorkplace = Workplace{
	SheetsTable: Sheets{
		SpreadsheetID: spreadsheetID,
		Range:         "TROBART!A:D",
	},

	WeeekProject: Weeek{
		Project:       Trobart,
		ProjectNumber: WeeekProjects[Trobart],
	},
}

var DripWorkplace = Workplace{
	SheetsTable: Sheets{
		SpreadsheetID: spreadsheetID,
		Range:         "DRIPS!A:D",
	},

	WeeekProject: Weeek{
		Project:       Drip,
		ProjectNumber: WeeekProjects[Drip],
	},
}

var CapsuleWorkplace = Workplace{
	SheetsTable: Sheets{
		SpreadsheetID: spreadsheetID,
		Range:         "CAPSULES!A:D",
	},

	WeeekProject: Weeek{
		Project:       Capsule,
		ProjectNumber: WeeekProjects[Capsule],
	},
}

var AssemblyWorkplace = Workplace{
	SheetsTable: Sheets{
		SpreadsheetID: spreadsheetID,
		Range:         "ASSEMBLY!A:D",
	},

	WeeekProject: Weeek{
		Project:       Assembly,
		ProjectNumber: WeeekProjects[Assembly],
	},
}

var Workplaces = []Workplace{
	Imf120Workplace,
	TrobartWorkplace,
	CapsuleWorkplace,
	DripWorkplace,
	AssemblyWorkplace,
}
