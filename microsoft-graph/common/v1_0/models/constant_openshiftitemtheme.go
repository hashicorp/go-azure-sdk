package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftItemTheme string

const (
	OpenShiftItemThemeblue       OpenShiftItemTheme = "Blue"
	OpenShiftItemThemedarkBlue   OpenShiftItemTheme = "DarkBlue"
	OpenShiftItemThemedarkGreen  OpenShiftItemTheme = "DarkGreen"
	OpenShiftItemThemedarkPink   OpenShiftItemTheme = "DarkPink"
	OpenShiftItemThemedarkPurple OpenShiftItemTheme = "DarkPurple"
	OpenShiftItemThemedarkYellow OpenShiftItemTheme = "DarkYellow"
	OpenShiftItemThemegray       OpenShiftItemTheme = "Gray"
	OpenShiftItemThemegreen      OpenShiftItemTheme = "Green"
	OpenShiftItemThemepink       OpenShiftItemTheme = "Pink"
	OpenShiftItemThemepurple     OpenShiftItemTheme = "Purple"
	OpenShiftItemThemewhite      OpenShiftItemTheme = "White"
	OpenShiftItemThemeyellow     OpenShiftItemTheme = "Yellow"
)

func PossibleValuesForOpenShiftItemTheme() []string {
	return []string{
		string(OpenShiftItemThemeblue),
		string(OpenShiftItemThemedarkBlue),
		string(OpenShiftItemThemedarkGreen),
		string(OpenShiftItemThemedarkPink),
		string(OpenShiftItemThemedarkPurple),
		string(OpenShiftItemThemedarkYellow),
		string(OpenShiftItemThemegray),
		string(OpenShiftItemThemegreen),
		string(OpenShiftItemThemepink),
		string(OpenShiftItemThemepurple),
		string(OpenShiftItemThemewhite),
		string(OpenShiftItemThemeyellow),
	}
}

func parseOpenShiftItemTheme(input string) (*OpenShiftItemTheme, error) {
	vals := map[string]OpenShiftItemTheme{
		"blue":       OpenShiftItemThemeblue,
		"darkblue":   OpenShiftItemThemedarkBlue,
		"darkgreen":  OpenShiftItemThemedarkGreen,
		"darkpink":   OpenShiftItemThemedarkPink,
		"darkpurple": OpenShiftItemThemedarkPurple,
		"darkyellow": OpenShiftItemThemedarkYellow,
		"gray":       OpenShiftItemThemegray,
		"green":      OpenShiftItemThemegreen,
		"pink":       OpenShiftItemThemepink,
		"purple":     OpenShiftItemThemepurple,
		"white":      OpenShiftItemThemewhite,
		"yellow":     OpenShiftItemThemeyellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenShiftItemTheme(input)
	return &out, nil
}
