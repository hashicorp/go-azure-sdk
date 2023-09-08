package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintOperationStatusState string

const (
	PrintOperationStatusStatefailed     PrintOperationStatusState = "Failed"
	PrintOperationStatusStatenotStarted PrintOperationStatusState = "NotStarted"
	PrintOperationStatusStaterunning    PrintOperationStatusState = "Running"
	PrintOperationStatusStatesucceeded  PrintOperationStatusState = "Succeeded"
)

func PossibleValuesForPrintOperationStatusState() []string {
	return []string{
		string(PrintOperationStatusStatefailed),
		string(PrintOperationStatusStatenotStarted),
		string(PrintOperationStatusStaterunning),
		string(PrintOperationStatusStatesucceeded),
	}
}

func parsePrintOperationStatusState(input string) (*PrintOperationStatusState, error) {
	vals := map[string]PrintOperationStatusState{
		"failed":     PrintOperationStatusStatefailed,
		"notstarted": PrintOperationStatusStatenotStarted,
		"running":    PrintOperationStatusStaterunning,
		"succeeded":  PrintOperationStatusStatesucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintOperationStatusState(input)
	return &out, nil
}
