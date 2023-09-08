package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarColor string

const (
	CalendarColorauto        CalendarColor = "Auto"
	CalendarColorlightBlue   CalendarColor = "LightBlue"
	CalendarColorlightBrown  CalendarColor = "LightBrown"
	CalendarColorlightGray   CalendarColor = "LightGray"
	CalendarColorlightGreen  CalendarColor = "LightGreen"
	CalendarColorlightOrange CalendarColor = "LightOrange"
	CalendarColorlightPink   CalendarColor = "LightPink"
	CalendarColorlightRed    CalendarColor = "LightRed"
	CalendarColorlightTeal   CalendarColor = "LightTeal"
	CalendarColorlightYellow CalendarColor = "LightYellow"
	CalendarColormaxColor    CalendarColor = "MaxColor"
)

func PossibleValuesForCalendarColor() []string {
	return []string{
		string(CalendarColorauto),
		string(CalendarColorlightBlue),
		string(CalendarColorlightBrown),
		string(CalendarColorlightGray),
		string(CalendarColorlightGreen),
		string(CalendarColorlightOrange),
		string(CalendarColorlightPink),
		string(CalendarColorlightRed),
		string(CalendarColorlightTeal),
		string(CalendarColorlightYellow),
		string(CalendarColormaxColor),
	}
}

func parseCalendarColor(input string) (*CalendarColor, error) {
	vals := map[string]CalendarColor{
		"auto":        CalendarColorauto,
		"lightblue":   CalendarColorlightBlue,
		"lightbrown":  CalendarColorlightBrown,
		"lightgray":   CalendarColorlightGray,
		"lightgreen":  CalendarColorlightGreen,
		"lightorange": CalendarColorlightOrange,
		"lightpink":   CalendarColorlightPink,
		"lightred":    CalendarColorlightRed,
		"lightteal":   CalendarColorlightTeal,
		"lightyellow": CalendarColorlightYellow,
		"maxcolor":    CalendarColormaxColor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarColor(input)
	return &out, nil
}
