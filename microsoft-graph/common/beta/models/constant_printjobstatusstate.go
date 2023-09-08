package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobStatusState string

const (
	PrintJobStatusStateaborted    PrintJobStatusState = "Aborted"
	PrintJobStatusStatecanceled   PrintJobStatusState = "Canceled"
	PrintJobStatusStatecompleted  PrintJobStatusState = "Completed"
	PrintJobStatusStatepaused     PrintJobStatusState = "Paused"
	PrintJobStatusStatepending    PrintJobStatusState = "Pending"
	PrintJobStatusStateprocessing PrintJobStatusState = "Processing"
	PrintJobStatusStatestopped    PrintJobStatusState = "Stopped"
	PrintJobStatusStateunknown    PrintJobStatusState = "Unknown"
)

func PossibleValuesForPrintJobStatusState() []string {
	return []string{
		string(PrintJobStatusStateaborted),
		string(PrintJobStatusStatecanceled),
		string(PrintJobStatusStatecompleted),
		string(PrintJobStatusStatepaused),
		string(PrintJobStatusStatepending),
		string(PrintJobStatusStateprocessing),
		string(PrintJobStatusStatestopped),
		string(PrintJobStatusStateunknown),
	}
}

func parsePrintJobStatusState(input string) (*PrintJobStatusState, error) {
	vals := map[string]PrintJobStatusState{
		"aborted":    PrintJobStatusStateaborted,
		"canceled":   PrintJobStatusStatecanceled,
		"completed":  PrintJobStatusStatecompleted,
		"paused":     PrintJobStatusStatepaused,
		"pending":    PrintJobStatusStatepending,
		"processing": PrintJobStatusStateprocessing,
		"stopped":    PrintJobStatusStatestopped,
		"unknown":    PrintJobStatusStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobStatusState(input)
	return &out, nil
}
