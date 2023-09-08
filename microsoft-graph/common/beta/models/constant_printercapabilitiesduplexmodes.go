package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesDuplexModes string

const (
	PrinterCapabilitiesDuplexModesflipOnLongEdge  PrinterCapabilitiesDuplexModes = "FlipOnLongEdge"
	PrinterCapabilitiesDuplexModesflipOnShortEdge PrinterCapabilitiesDuplexModes = "FlipOnShortEdge"
	PrinterCapabilitiesDuplexModesoneSided        PrinterCapabilitiesDuplexModes = "OneSided"
)

func PossibleValuesForPrinterCapabilitiesDuplexModes() []string {
	return []string{
		string(PrinterCapabilitiesDuplexModesflipOnLongEdge),
		string(PrinterCapabilitiesDuplexModesflipOnShortEdge),
		string(PrinterCapabilitiesDuplexModesoneSided),
	}
}

func parsePrinterCapabilitiesDuplexModes(input string) (*PrinterCapabilitiesDuplexModes, error) {
	vals := map[string]PrinterCapabilitiesDuplexModes{
		"fliponlongedge":  PrinterCapabilitiesDuplexModesflipOnLongEdge,
		"fliponshortedge": PrinterCapabilitiesDuplexModesflipOnShortEdge,
		"onesided":        PrinterCapabilitiesDuplexModesoneSided,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesDuplexModes(input)
	return &out, nil
}
