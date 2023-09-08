package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleEntityTheme string

const (
	ScheduleEntityThemeblue       ScheduleEntityTheme = "Blue"
	ScheduleEntityThemedarkBlue   ScheduleEntityTheme = "DarkBlue"
	ScheduleEntityThemedarkGreen  ScheduleEntityTheme = "DarkGreen"
	ScheduleEntityThemedarkPink   ScheduleEntityTheme = "DarkPink"
	ScheduleEntityThemedarkPurple ScheduleEntityTheme = "DarkPurple"
	ScheduleEntityThemedarkYellow ScheduleEntityTheme = "DarkYellow"
	ScheduleEntityThemegray       ScheduleEntityTheme = "Gray"
	ScheduleEntityThemegreen      ScheduleEntityTheme = "Green"
	ScheduleEntityThemepink       ScheduleEntityTheme = "Pink"
	ScheduleEntityThemepurple     ScheduleEntityTheme = "Purple"
	ScheduleEntityThemewhite      ScheduleEntityTheme = "White"
	ScheduleEntityThemeyellow     ScheduleEntityTheme = "Yellow"
)

func PossibleValuesForScheduleEntityTheme() []string {
	return []string{
		string(ScheduleEntityThemeblue),
		string(ScheduleEntityThemedarkBlue),
		string(ScheduleEntityThemedarkGreen),
		string(ScheduleEntityThemedarkPink),
		string(ScheduleEntityThemedarkPurple),
		string(ScheduleEntityThemedarkYellow),
		string(ScheduleEntityThemegray),
		string(ScheduleEntityThemegreen),
		string(ScheduleEntityThemepink),
		string(ScheduleEntityThemepurple),
		string(ScheduleEntityThemewhite),
		string(ScheduleEntityThemeyellow),
	}
}

func parseScheduleEntityTheme(input string) (*ScheduleEntityTheme, error) {
	vals := map[string]ScheduleEntityTheme{
		"blue":       ScheduleEntityThemeblue,
		"darkblue":   ScheduleEntityThemedarkBlue,
		"darkgreen":  ScheduleEntityThemedarkGreen,
		"darkpink":   ScheduleEntityThemedarkPink,
		"darkpurple": ScheduleEntityThemedarkPurple,
		"darkyellow": ScheduleEntityThemedarkYellow,
		"gray":       ScheduleEntityThemegray,
		"green":      ScheduleEntityThemegreen,
		"pink":       ScheduleEntityThemepink,
		"purple":     ScheduleEntityThemepurple,
		"white":      ScheduleEntityThemewhite,
		"yellow":     ScheduleEntityThemeyellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleEntityTheme(input)
	return &out, nil
}
