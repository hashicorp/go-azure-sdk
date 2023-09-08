package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsPrintQuality string

const (
	PrinterDefaultsPrintQualityhigh   PrinterDefaultsPrintQuality = "High"
	PrinterDefaultsPrintQualitylow    PrinterDefaultsPrintQuality = "Low"
	PrinterDefaultsPrintQualitymedium PrinterDefaultsPrintQuality = "Medium"
)

func PossibleValuesForPrinterDefaultsPrintQuality() []string {
	return []string{
		string(PrinterDefaultsPrintQualityhigh),
		string(PrinterDefaultsPrintQualitylow),
		string(PrinterDefaultsPrintQualitymedium),
	}
}

func parsePrinterDefaultsPrintQuality(input string) (*PrinterDefaultsPrintQuality, error) {
	vals := map[string]PrinterDefaultsPrintQuality{
		"high":   PrinterDefaultsPrintQualityhigh,
		"low":    PrinterDefaultsPrintQualitylow,
		"medium": PrinterDefaultsPrintQualitymedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsPrintQuality(input)
	return &out, nil
}
