package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsDuplexConfiguration string

const (
	PrinterDefaultsDuplexConfigurationoneSided          PrinterDefaultsDuplexConfiguration = "OneSided"
	PrinterDefaultsDuplexConfigurationtwoSidedLongEdge  PrinterDefaultsDuplexConfiguration = "TwoSidedLongEdge"
	PrinterDefaultsDuplexConfigurationtwoSidedShortEdge PrinterDefaultsDuplexConfiguration = "TwoSidedShortEdge"
)

func PossibleValuesForPrinterDefaultsDuplexConfiguration() []string {
	return []string{
		string(PrinterDefaultsDuplexConfigurationoneSided),
		string(PrinterDefaultsDuplexConfigurationtwoSidedLongEdge),
		string(PrinterDefaultsDuplexConfigurationtwoSidedShortEdge),
	}
}

func parsePrinterDefaultsDuplexConfiguration(input string) (*PrinterDefaultsDuplexConfiguration, error) {
	vals := map[string]PrinterDefaultsDuplexConfiguration{
		"onesided":          PrinterDefaultsDuplexConfigurationoneSided,
		"twosidedlongedge":  PrinterDefaultsDuplexConfigurationtwoSidedLongEdge,
		"twosidedshortedge": PrinterDefaultsDuplexConfigurationtwoSidedShortEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsDuplexConfiguration(input)
	return &out, nil
}
