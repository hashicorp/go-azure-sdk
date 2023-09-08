package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterStatusState string

const (
	PrinterStatusStateidle       PrinterStatusState = "Idle"
	PrinterStatusStateprocessing PrinterStatusState = "Processing"
	PrinterStatusStatestopped    PrinterStatusState = "Stopped"
	PrinterStatusStateunknown    PrinterStatusState = "Unknown"
)

func PossibleValuesForPrinterStatusState() []string {
	return []string{
		string(PrinterStatusStateidle),
		string(PrinterStatusStateprocessing),
		string(PrinterStatusStatestopped),
		string(PrinterStatusStateunknown),
	}
}

func parsePrinterStatusState(input string) (*PrinterStatusState, error) {
	vals := map[string]PrinterStatusState{
		"idle":       PrinterStatusStateidle,
		"processing": PrinterStatusStateprocessing,
		"stopped":    PrinterStatusStatestopped,
		"unknown":    PrinterStatusStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterStatusState(input)
	return &out, nil
}
