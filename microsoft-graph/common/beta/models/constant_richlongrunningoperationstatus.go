package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RichLongRunningOperationStatus string

const (
	RichLongRunningOperationStatusfailed     RichLongRunningOperationStatus = "Failed"
	RichLongRunningOperationStatusnotStarted RichLongRunningOperationStatus = "NotStarted"
	RichLongRunningOperationStatusrunning    RichLongRunningOperationStatus = "Running"
	RichLongRunningOperationStatussucceeded  RichLongRunningOperationStatus = "Succeeded"
)

func PossibleValuesForRichLongRunningOperationStatus() []string {
	return []string{
		string(RichLongRunningOperationStatusfailed),
		string(RichLongRunningOperationStatusnotStarted),
		string(RichLongRunningOperationStatusrunning),
		string(RichLongRunningOperationStatussucceeded),
	}
}

func parseRichLongRunningOperationStatus(input string) (*RichLongRunningOperationStatus, error) {
	vals := map[string]RichLongRunningOperationStatus{
		"failed":     RichLongRunningOperationStatusfailed,
		"notstarted": RichLongRunningOperationStatusnotStarted,
		"running":    RichLongRunningOperationStatusrunning,
		"succeeded":  RichLongRunningOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RichLongRunningOperationStatus(input)
	return &out, nil
}
