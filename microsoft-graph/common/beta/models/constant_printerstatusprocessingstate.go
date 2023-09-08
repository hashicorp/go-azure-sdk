package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterStatusProcessingState string

const (
	PrinterStatusProcessingStateidle       PrinterStatusProcessingState = "Idle"
	PrinterStatusProcessingStateprocessing PrinterStatusProcessingState = "Processing"
	PrinterStatusProcessingStatestopped    PrinterStatusProcessingState = "Stopped"
	PrinterStatusProcessingStateunknown    PrinterStatusProcessingState = "Unknown"
)

func PossibleValuesForPrinterStatusProcessingState() []string {
	return []string{
		string(PrinterStatusProcessingStateidle),
		string(PrinterStatusProcessingStateprocessing),
		string(PrinterStatusProcessingStatestopped),
		string(PrinterStatusProcessingStateunknown),
	}
}

func parsePrinterStatusProcessingState(input string) (*PrinterStatusProcessingState, error) {
	vals := map[string]PrinterStatusProcessingState{
		"idle":       PrinterStatusProcessingStateidle,
		"processing": PrinterStatusProcessingStateprocessing,
		"stopped":    PrinterStatusProcessingStatestopped,
		"unknown":    PrinterStatusProcessingStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterStatusProcessingState(input)
	return &out, nil
}
