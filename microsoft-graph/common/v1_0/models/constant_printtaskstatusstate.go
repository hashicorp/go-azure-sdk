package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintTaskStatusState string

const (
	PrintTaskStatusStateaborted    PrintTaskStatusState = "Aborted"
	PrintTaskStatusStatecompleted  PrintTaskStatusState = "Completed"
	PrintTaskStatusStatepending    PrintTaskStatusState = "Pending"
	PrintTaskStatusStateprocessing PrintTaskStatusState = "Processing"
)

func PossibleValuesForPrintTaskStatusState() []string {
	return []string{
		string(PrintTaskStatusStateaborted),
		string(PrintTaskStatusStatecompleted),
		string(PrintTaskStatusStatepending),
		string(PrintTaskStatusStateprocessing),
	}
}

func parsePrintTaskStatusState(input string) (*PrintTaskStatusState, error) {
	vals := map[string]PrintTaskStatusState{
		"aborted":    PrintTaskStatusStateaborted,
		"completed":  PrintTaskStatusStatecompleted,
		"pending":    PrintTaskStatusStatepending,
		"processing": PrintTaskStatusStateprocessing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintTaskStatusState(input)
	return &out, nil
}
