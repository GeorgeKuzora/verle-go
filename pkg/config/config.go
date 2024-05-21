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
		UpdateRange:   "IMF!A1:D200",
		SheetID:       1330258137,
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
		UpdateRange:   "TROBART!A1:D200",
		SheetID:       1738797376,
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
		UpdateRange:   "DRIPS!A1:D200",
		SheetID:       612152640,
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
		UpdateRange:   "CAPSULES!A1:D200",
		SheetID:       1199560039,
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
		UpdateRange:   "ASSEMBLY!A1:D200",
		SheetID:       1355663488,
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
