package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobStatusProcessingState string

const (
	PrintJobStatusProcessingStateaborted    PrintJobStatusProcessingState = "Aborted"
	PrintJobStatusProcessingStatecanceled   PrintJobStatusProcessingState = "Canceled"
	PrintJobStatusProcessingStatecompleted  PrintJobStatusProcessingState = "Completed"
	PrintJobStatusProcessingStatepaused     PrintJobStatusProcessingState = "Paused"
	PrintJobStatusProcessingStatepending    PrintJobStatusProcessingState = "Pending"
	PrintJobStatusProcessingStateprocessing PrintJobStatusProcessingState = "Processing"
	PrintJobStatusProcessingStatestopped    PrintJobStatusProcessingState = "Stopped"
	PrintJobStatusProcessingStateunknown    PrintJobStatusProcessingState = "Unknown"
)

func PossibleValuesForPrintJobStatusProcessingState() []string {
	return []string{
		string(PrintJobStatusProcessingStateaborted),
		string(PrintJobStatusProcessingStatecanceled),
		string(PrintJobStatusProcessingStatecompleted),
		string(PrintJobStatusProcessingStatepaused),
		string(PrintJobStatusProcessingStatepending),
		string(PrintJobStatusProcessingStateprocessing),
		string(PrintJobStatusProcessingStatestopped),
		string(PrintJobStatusProcessingStateunknown),
	}
}

func parsePrintJobStatusProcessingState(input string) (*PrintJobStatusProcessingState, error) {
	vals := map[string]PrintJobStatusProcessingState{
		"aborted":    PrintJobStatusProcessingStateaborted,
		"canceled":   PrintJobStatusProcessingStatecanceled,
		"completed":  PrintJobStatusProcessingStatecompleted,
		"paused":     PrintJobStatusProcessingStatepaused,
		"pending":    PrintJobStatusProcessingStatepending,
		"processing": PrintJobStatusProcessingStateprocessing,
		"stopped":    PrintJobStatusProcessingStatestopped,
		"unknown":    PrintJobStatusProcessingStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobStatusProcessingState(input)
	return &out, nil
}
