package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShiftItemTheme string

const (
	ShiftItemThemeblue       ShiftItemTheme = "Blue"
	ShiftItemThemedarkBlue   ShiftItemTheme = "DarkBlue"
	ShiftItemThemedarkGreen  ShiftItemTheme = "DarkGreen"
	ShiftItemThemedarkPink   ShiftItemTheme = "DarkPink"
	ShiftItemThemedarkPurple ShiftItemTheme = "DarkPurple"
	ShiftItemThemedarkYellow ShiftItemTheme = "DarkYellow"
	ShiftItemThemegray       ShiftItemTheme = "Gray"
	ShiftItemThemegreen      ShiftItemTheme = "Green"
	ShiftItemThemepink       ShiftItemTheme = "Pink"
	ShiftItemThemepurple     ShiftItemTheme = "Purple"
	ShiftItemThemewhite      ShiftItemTheme = "White"
	ShiftItemThemeyellow     ShiftItemTheme = "Yellow"
)

func PossibleValuesForShiftItemTheme() []string {
	return []string{
		string(ShiftItemThemeblue),
		string(ShiftItemThemedarkBlue),
		string(ShiftItemThemedarkGreen),
		string(ShiftItemThemedarkPink),
		string(ShiftItemThemedarkPurple),
		string(ShiftItemThemedarkYellow),
		string(ShiftItemThemegray),
		string(ShiftItemThemegreen),
		string(ShiftItemThemepink),
		string(ShiftItemThemepurple),
		string(ShiftItemThemewhite),
		string(ShiftItemThemeyellow),
	}
}

func parseShiftItemTheme(input string) (*ShiftItemTheme, error) {
	vals := map[string]ShiftItemTheme{
		"blue":       ShiftItemThemeblue,
		"darkblue":   ShiftItemThemedarkBlue,
		"darkgreen":  ShiftItemThemedarkGreen,
		"darkpink":   ShiftItemThemedarkPink,
		"darkpurple": ShiftItemThemedarkPurple,
		"darkyellow": ShiftItemThemedarkYellow,
		"gray":       ShiftItemThemegray,
		"green":      ShiftItemThemegreen,
		"pink":       ShiftItemThemepink,
		"purple":     ShiftItemThemepurple,
		"white":      ShiftItemThemewhite,
		"yellow":     ShiftItemThemeyellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShiftItemTheme(input)
	return &out, nil
}
