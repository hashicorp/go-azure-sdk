package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShiftActivityTheme string

const (
	ShiftActivityThemeblue       ShiftActivityTheme = "Blue"
	ShiftActivityThemedarkBlue   ShiftActivityTheme = "DarkBlue"
	ShiftActivityThemedarkGreen  ShiftActivityTheme = "DarkGreen"
	ShiftActivityThemedarkPink   ShiftActivityTheme = "DarkPink"
	ShiftActivityThemedarkPurple ShiftActivityTheme = "DarkPurple"
	ShiftActivityThemedarkYellow ShiftActivityTheme = "DarkYellow"
	ShiftActivityThemegray       ShiftActivityTheme = "Gray"
	ShiftActivityThemegreen      ShiftActivityTheme = "Green"
	ShiftActivityThemepink       ShiftActivityTheme = "Pink"
	ShiftActivityThemepurple     ShiftActivityTheme = "Purple"
	ShiftActivityThemewhite      ShiftActivityTheme = "White"
	ShiftActivityThemeyellow     ShiftActivityTheme = "Yellow"
)

func PossibleValuesForShiftActivityTheme() []string {
	return []string{
		string(ShiftActivityThemeblue),
		string(ShiftActivityThemedarkBlue),
		string(ShiftActivityThemedarkGreen),
		string(ShiftActivityThemedarkPink),
		string(ShiftActivityThemedarkPurple),
		string(ShiftActivityThemedarkYellow),
		string(ShiftActivityThemegray),
		string(ShiftActivityThemegreen),
		string(ShiftActivityThemepink),
		string(ShiftActivityThemepurple),
		string(ShiftActivityThemewhite),
		string(ShiftActivityThemeyellow),
	}
}

func parseShiftActivityTheme(input string) (*ShiftActivityTheme, error) {
	vals := map[string]ShiftActivityTheme{
		"blue":       ShiftActivityThemeblue,
		"darkblue":   ShiftActivityThemedarkBlue,
		"darkgreen":  ShiftActivityThemedarkGreen,
		"darkpink":   ShiftActivityThemedarkPink,
		"darkpurple": ShiftActivityThemedarkPurple,
		"darkyellow": ShiftActivityThemedarkYellow,
		"gray":       ShiftActivityThemegray,
		"green":      ShiftActivityThemegreen,
		"pink":       ShiftActivityThemepink,
		"purple":     ShiftActivityThemepurple,
		"white":      ShiftActivityThemewhite,
		"yellow":     ShiftActivityThemeyellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShiftActivityTheme(input)
	return &out, nil
}
