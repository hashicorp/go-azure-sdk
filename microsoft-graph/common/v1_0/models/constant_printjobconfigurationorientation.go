package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationOrientation string

const (
	PrintJobConfigurationOrientationlandscape        PrintJobConfigurationOrientation = "Landscape"
	PrintJobConfigurationOrientationportrait         PrintJobConfigurationOrientation = "Portrait"
	PrintJobConfigurationOrientationreverseLandscape PrintJobConfigurationOrientation = "ReverseLandscape"
	PrintJobConfigurationOrientationreversePortrait  PrintJobConfigurationOrientation = "ReversePortrait"
)

func PossibleValuesForPrintJobConfigurationOrientation() []string {
	return []string{
		string(PrintJobConfigurationOrientationlandscape),
		string(PrintJobConfigurationOrientationportrait),
		string(PrintJobConfigurationOrientationreverseLandscape),
		string(PrintJobConfigurationOrientationreversePortrait),
	}
}

func parsePrintJobConfigurationOrientation(input string) (*PrintJobConfigurationOrientation, error) {
	vals := map[string]PrintJobConfigurationOrientation{
		"landscape":        PrintJobConfigurationOrientationlandscape,
		"portrait":         PrintJobConfigurationOrientationportrait,
		"reverselandscape": PrintJobConfigurationOrientationreverseLandscape,
		"reverseportrait":  PrintJobConfigurationOrientationreversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationOrientation(input)
	return &out, nil
}
