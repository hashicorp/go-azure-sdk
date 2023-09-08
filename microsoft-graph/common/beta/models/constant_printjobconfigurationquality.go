package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationQuality string

const (
	PrintJobConfigurationQualityhigh   PrintJobConfigurationQuality = "High"
	PrintJobConfigurationQualitylow    PrintJobConfigurationQuality = "Low"
	PrintJobConfigurationQualitymedium PrintJobConfigurationQuality = "Medium"
)

func PossibleValuesForPrintJobConfigurationQuality() []string {
	return []string{
		string(PrintJobConfigurationQualityhigh),
		string(PrintJobConfigurationQualitylow),
		string(PrintJobConfigurationQualitymedium),
	}
}

func parsePrintJobConfigurationQuality(input string) (*PrintJobConfigurationQuality, error) {
	vals := map[string]PrintJobConfigurationQuality{
		"high":   PrintJobConfigurationQualityhigh,
		"low":    PrintJobConfigurationQualitylow,
		"medium": PrintJobConfigurationQualitymedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationQuality(input)
	return &out, nil
}
