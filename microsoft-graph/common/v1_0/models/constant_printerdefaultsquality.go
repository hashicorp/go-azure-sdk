package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsQuality string

const (
	PrinterDefaultsQualityhigh   PrinterDefaultsQuality = "High"
	PrinterDefaultsQualitylow    PrinterDefaultsQuality = "Low"
	PrinterDefaultsQualitymedium PrinterDefaultsQuality = "Medium"
)

func PossibleValuesForPrinterDefaultsQuality() []string {
	return []string{
		string(PrinterDefaultsQualityhigh),
		string(PrinterDefaultsQualitylow),
		string(PrinterDefaultsQualitymedium),
	}
}

func parsePrinterDefaultsQuality(input string) (*PrinterDefaultsQuality, error) {
	vals := map[string]PrinterDefaultsQuality{
		"high":   PrinterDefaultsQualityhigh,
		"low":    PrinterDefaultsQualitylow,
		"medium": PrinterDefaultsQualitymedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsQuality(input)
	return &out, nil
}
