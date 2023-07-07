package longtermretentionbackup

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExecutionStatus string

const (
	ExecutionStatusCancelled ExecutionStatus = "Cancelled"
	ExecutionStatusFailed    ExecutionStatus = "Failed"
	ExecutionStatusRunning   ExecutionStatus = "Running"
	ExecutionStatusSucceeded ExecutionStatus = "Succeeded"
)

func PossibleValuesForExecutionStatus() []string {
	return []string{
		string(ExecutionStatusCancelled),
		string(ExecutionStatusFailed),
		string(ExecutionStatusRunning),
		string(ExecutionStatusSucceeded),
	}
}

func parseExecutionStatus(input string) (*ExecutionStatus, error) {
	vals := map[string]ExecutionStatus{
		"cancelled": ExecutionStatusCancelled,
		"failed":    ExecutionStatusFailed,
		"running":   ExecutionStatusRunning,
		"succeeded": ExecutionStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExecutionStatus(input)
	return &out, nil
}
