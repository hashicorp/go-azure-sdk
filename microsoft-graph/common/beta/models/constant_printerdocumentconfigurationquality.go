package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationQuality string

const (
	PrinterDocumentConfigurationQualityhigh   PrinterDocumentConfigurationQuality = "High"
	PrinterDocumentConfigurationQualitylow    PrinterDocumentConfigurationQuality = "Low"
	PrinterDocumentConfigurationQualitymedium PrinterDocumentConfigurationQuality = "Medium"
)

func PossibleValuesForPrinterDocumentConfigurationQuality() []string {
	return []string{
		string(PrinterDocumentConfigurationQualityhigh),
		string(PrinterDocumentConfigurationQualitylow),
		string(PrinterDocumentConfigurationQualitymedium),
	}
}

func parsePrinterDocumentConfigurationQuality(input string) (*PrinterDocumentConfigurationQuality, error) {
	vals := map[string]PrinterDocumentConfigurationQuality{
		"high":   PrinterDocumentConfigurationQualityhigh,
		"low":    PrinterDocumentConfigurationQualitylow,
		"medium": PrinterDocumentConfigurationQualitymedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationQuality(input)
	return &out, nil
}
