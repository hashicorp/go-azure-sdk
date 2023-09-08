package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedPrintQualities string

const (
	PrinterCapabilitiesSupportedPrintQualitieshigh   PrinterCapabilitiesSupportedPrintQualities = "High"
	PrinterCapabilitiesSupportedPrintQualitieslow    PrinterCapabilitiesSupportedPrintQualities = "Low"
	PrinterCapabilitiesSupportedPrintQualitiesmedium PrinterCapabilitiesSupportedPrintQualities = "Medium"
)

func PossibleValuesForPrinterCapabilitiesSupportedPrintQualities() []string {
	return []string{
		string(PrinterCapabilitiesSupportedPrintQualitieshigh),
		string(PrinterCapabilitiesSupportedPrintQualitieslow),
		string(PrinterCapabilitiesSupportedPrintQualitiesmedium),
	}
}

func parsePrinterCapabilitiesSupportedPrintQualities(input string) (*PrinterCapabilitiesSupportedPrintQualities, error) {
	vals := map[string]PrinterCapabilitiesSupportedPrintQualities{
		"high":   PrinterCapabilitiesSupportedPrintQualitieshigh,
		"low":    PrinterCapabilitiesSupportedPrintQualitieslow,
		"medium": PrinterCapabilitiesSupportedPrintQualitiesmedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedPrintQualities(input)
	return &out, nil
}
