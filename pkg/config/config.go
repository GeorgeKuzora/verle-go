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
		Credentials:   credentials,
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
		Credentials:   credentials,
	},

	WeeekProject: Weeek{
		Project:       Trobart,
		ProjectNumber: WeeekProjects[Trobart],
	},
}

var DripWorkplace = Workplace{
	SheetsTable: Sheets{
		SpreadsheetID: spreadsheetID,
		Range:         "DRIP!A:D",
		Credentials:   credentials,
	},

	WeeekProject: Weeek{
		Project:       Drip,
		ProjectNumber: WeeekProjects[Drip],
	},
}

var CapsuleWorkplace = Workplace{
	SheetsTable: Sheets{
		SpreadsheetID: spreadsheetID,
		Range:         "CAPSULE!A:D",
		Credentials:   credentials,
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
		Credentials:   credentials,
	},

	WeeekProject: Weeek{
		Project:       Assembly,
		ProjectNumber: WeeekProjects[Assembly],
	},
}
