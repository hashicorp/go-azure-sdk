package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationDuplexMode string

const (
	PrinterDocumentConfigurationDuplexModeflipOnLongEdge  PrinterDocumentConfigurationDuplexMode = "FlipOnLongEdge"
	PrinterDocumentConfigurationDuplexModeflipOnShortEdge PrinterDocumentConfigurationDuplexMode = "FlipOnShortEdge"
	PrinterDocumentConfigurationDuplexModeoneSided        PrinterDocumentConfigurationDuplexMode = "OneSided"
)

func PossibleValuesForPrinterDocumentConfigurationDuplexMode() []string {
	return []string{
		string(PrinterDocumentConfigurationDuplexModeflipOnLongEdge),
		string(PrinterDocumentConfigurationDuplexModeflipOnShortEdge),
		string(PrinterDocumentConfigurationDuplexModeoneSided),
	}
}

func parsePrinterDocumentConfigurationDuplexMode(input string) (*PrinterDocumentConfigurationDuplexMode, error) {
	vals := map[string]PrinterDocumentConfigurationDuplexMode{
		"fliponlongedge":  PrinterDocumentConfigurationDuplexModeflipOnLongEdge,
		"fliponshortedge": PrinterDocumentConfigurationDuplexModeflipOnShortEdge,
		"onesided":        PrinterDocumentConfigurationDuplexModeoneSided,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationDuplexMode(input)
	return &out, nil
}
