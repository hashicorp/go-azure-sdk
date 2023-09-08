package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsOrientation string

const (
	PrinterDefaultsOrientationlandscape        PrinterDefaultsOrientation = "Landscape"
	PrinterDefaultsOrientationportrait         PrinterDefaultsOrientation = "Portrait"
	PrinterDefaultsOrientationreverseLandscape PrinterDefaultsOrientation = "ReverseLandscape"
	PrinterDefaultsOrientationreversePortrait  PrinterDefaultsOrientation = "ReversePortrait"
)

func PossibleValuesForPrinterDefaultsOrientation() []string {
	return []string{
		string(PrinterDefaultsOrientationlandscape),
		string(PrinterDefaultsOrientationportrait),
		string(PrinterDefaultsOrientationreverseLandscape),
		string(PrinterDefaultsOrientationreversePortrait),
	}
}

func parsePrinterDefaultsOrientation(input string) (*PrinterDefaultsOrientation, error) {
	vals := map[string]PrinterDefaultsOrientation{
		"landscape":        PrinterDefaultsOrientationlandscape,
		"portrait":         PrinterDefaultsOrientationportrait,
		"reverselandscape": PrinterDefaultsOrientationreverseLandscape,
		"reverseportrait":  PrinterDefaultsOrientationreversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsOrientation(input)
	return &out, nil
}
