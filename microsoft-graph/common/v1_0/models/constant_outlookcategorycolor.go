package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookCategoryColor string

const (
	OutlookCategoryColornone     OutlookCategoryColor = "None"
	OutlookCategoryColorpreset0  OutlookCategoryColor = "Preset0"
	OutlookCategoryColorpreset1  OutlookCategoryColor = "Preset1"
	OutlookCategoryColorpreset10 OutlookCategoryColor = "Preset10"
	OutlookCategoryColorpreset11 OutlookCategoryColor = "Preset11"
	OutlookCategoryColorpreset12 OutlookCategoryColor = "Preset12"
	OutlookCategoryColorpreset13 OutlookCategoryColor = "Preset13"
	OutlookCategoryColorpreset14 OutlookCategoryColor = "Preset14"
	OutlookCategoryColorpreset15 OutlookCategoryColor = "Preset15"
	OutlookCategoryColorpreset16 OutlookCategoryColor = "Preset16"
	OutlookCategoryColorpreset17 OutlookCategoryColor = "Preset17"
	OutlookCategoryColorpreset18 OutlookCategoryColor = "Preset18"
	OutlookCategoryColorpreset19 OutlookCategoryColor = "Preset19"
	OutlookCategoryColorpreset2  OutlookCategoryColor = "Preset2"
	OutlookCategoryColorpreset20 OutlookCategoryColor = "Preset20"
	OutlookCategoryColorpreset21 OutlookCategoryColor = "Preset21"
	OutlookCategoryColorpreset22 OutlookCategoryColor = "Preset22"
	OutlookCategoryColorpreset23 OutlookCategoryColor = "Preset23"
	OutlookCategoryColorpreset24 OutlookCategoryColor = "Preset24"
	OutlookCategoryColorpreset3  OutlookCategoryColor = "Preset3"
	OutlookCategoryColorpreset4  OutlookCategoryColor = "Preset4"
	OutlookCategoryColorpreset5  OutlookCategoryColor = "Preset5"
	OutlookCategoryColorpreset6  OutlookCategoryColor = "Preset6"
	OutlookCategoryColorpreset7  OutlookCategoryColor = "Preset7"
	OutlookCategoryColorpreset8  OutlookCategoryColor = "Preset8"
	OutlookCategoryColorpreset9  OutlookCategoryColor = "Preset9"
)

func PossibleValuesForOutlookCategoryColor() []string {
	return []string{
		string(OutlookCategoryColornone),
		string(OutlookCategoryColorpreset0),
		string(OutlookCategoryColorpreset1),
		string(OutlookCategoryColorpreset10),
		string(OutlookCategoryColorpreset11),
		string(OutlookCategoryColorpreset12),
		string(OutlookCategoryColorpreset13),
		string(OutlookCategoryColorpreset14),
		string(OutlookCategoryColorpreset15),
		string(OutlookCategoryColorpreset16),
		string(OutlookCategoryColorpreset17),
		string(OutlookCategoryColorpreset18),
		string(OutlookCategoryColorpreset19),
		string(OutlookCategoryColorpreset2),
		string(OutlookCategoryColorpreset20),
		string(OutlookCategoryColorpreset21),
		string(OutlookCategoryColorpreset22),
		string(OutlookCategoryColorpreset23),
		string(OutlookCategoryColorpreset24),
		string(OutlookCategoryColorpreset3),
		string(OutlookCategoryColorpreset4),
		string(OutlookCategoryColorpreset5),
		string(OutlookCategoryColorpreset6),
		string(OutlookCategoryColorpreset7),
		string(OutlookCategoryColorpreset8),
		string(OutlookCategoryColorpreset9),
	}
}

func parseOutlookCategoryColor(input string) (*OutlookCategoryColor, error) {
	vals := map[string]OutlookCategoryColor{
		"none":     OutlookCategoryColornone,
		"preset0":  OutlookCategoryColorpreset0,
		"preset1":  OutlookCategoryColorpreset1,
		"preset10": OutlookCategoryColorpreset10,
		"preset11": OutlookCategoryColorpreset11,
		"preset12": OutlookCategoryColorpreset12,
		"preset13": OutlookCategoryColorpreset13,
		"preset14": OutlookCategoryColorpreset14,
		"preset15": OutlookCategoryColorpreset15,
		"preset16": OutlookCategoryColorpreset16,
		"preset17": OutlookCategoryColorpreset17,
		"preset18": OutlookCategoryColorpreset18,
		"preset19": OutlookCategoryColorpreset19,
		"preset2":  OutlookCategoryColorpreset2,
		"preset20": OutlookCategoryColorpreset20,
		"preset21": OutlookCategoryColorpreset21,
		"preset22": OutlookCategoryColorpreset22,
		"preset23": OutlookCategoryColorpreset23,
		"preset24": OutlookCategoryColorpreset24,
		"preset3":  OutlookCategoryColorpreset3,
		"preset4":  OutlookCategoryColorpreset4,
		"preset5":  OutlookCategoryColorpreset5,
		"preset6":  OutlookCategoryColorpreset6,
		"preset7":  OutlookCategoryColorpreset7,
		"preset8":  OutlookCategoryColorpreset8,
		"preset9":  OutlookCategoryColorpreset9,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlookCategoryColor(input)
	return &out, nil
}
