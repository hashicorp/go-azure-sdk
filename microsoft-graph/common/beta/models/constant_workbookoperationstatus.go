package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookOperationStatus string

const (
	WorkbookOperationStatusfailed     WorkbookOperationStatus = "Failed"
	WorkbookOperationStatusnotStarted WorkbookOperationStatus = "NotStarted"
	WorkbookOperationStatusrunning    WorkbookOperationStatus = "Running"
	WorkbookOperationStatussucceeded  WorkbookOperationStatus = "Succeeded"
)

func PossibleValuesForWorkbookOperationStatus() []string {
	return []string{
		string(WorkbookOperationStatusfailed),
		string(WorkbookOperationStatusnotStarted),
		string(WorkbookOperationStatusrunning),
		string(WorkbookOperationStatussucceeded),
	}
}

func parseWorkbookOperationStatus(input string) (*WorkbookOperationStatus, error) {
	vals := map[string]WorkbookOperationStatus{
		"failed":     WorkbookOperationStatusfailed,
		"notstarted": WorkbookOperationStatusnotStarted,
		"running":    WorkbookOperationStatusrunning,
		"succeeded":  WorkbookOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkbookOperationStatus(input)
	return &out, nil
}
