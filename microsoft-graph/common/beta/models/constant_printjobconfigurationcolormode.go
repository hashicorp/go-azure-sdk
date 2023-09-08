package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationColorMode string

const (
	PrintJobConfigurationColorModeauto          PrintJobConfigurationColorMode = "Auto"
	PrintJobConfigurationColorModeblackAndWhite PrintJobConfigurationColorMode = "BlackAndWhite"
	PrintJobConfigurationColorModecolor         PrintJobConfigurationColorMode = "Color"
	PrintJobConfigurationColorModegrayscale     PrintJobConfigurationColorMode = "Grayscale"
)

func PossibleValuesForPrintJobConfigurationColorMode() []string {
	return []string{
		string(PrintJobConfigurationColorModeauto),
		string(PrintJobConfigurationColorModeblackAndWhite),
		string(PrintJobConfigurationColorModecolor),
		string(PrintJobConfigurationColorModegrayscale),
	}
}

func parsePrintJobConfigurationColorMode(input string) (*PrintJobConfigurationColorMode, error) {
	vals := map[string]PrintJobConfigurationColorMode{
		"auto":          PrintJobConfigurationColorModeauto,
		"blackandwhite": PrintJobConfigurationColorModeblackAndWhite,
		"color":         PrintJobConfigurationColorModecolor,
		"grayscale":     PrintJobConfigurationColorModegrayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationColorMode(input)
	return &out, nil
}
