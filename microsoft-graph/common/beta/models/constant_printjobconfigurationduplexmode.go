package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationDuplexMode string

const (
	PrintJobConfigurationDuplexModeflipOnLongEdge  PrintJobConfigurationDuplexMode = "FlipOnLongEdge"
	PrintJobConfigurationDuplexModeflipOnShortEdge PrintJobConfigurationDuplexMode = "FlipOnShortEdge"
	PrintJobConfigurationDuplexModeoneSided        PrintJobConfigurationDuplexMode = "OneSided"
)

func PossibleValuesForPrintJobConfigurationDuplexMode() []string {
	return []string{
		string(PrintJobConfigurationDuplexModeflipOnLongEdge),
		string(PrintJobConfigurationDuplexModeflipOnShortEdge),
		string(PrintJobConfigurationDuplexModeoneSided),
	}
}

func parsePrintJobConfigurationDuplexMode(input string) (*PrintJobConfigurationDuplexMode, error) {
	vals := map[string]PrintJobConfigurationDuplexMode{
		"fliponlongedge":  PrintJobConfigurationDuplexModeflipOnLongEdge,
		"fliponshortedge": PrintJobConfigurationDuplexModeflipOnShortEdge,
		"onesided":        PrintJobConfigurationDuplexModeoneSided,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationDuplexMode(input)
	return &out, nil
}
