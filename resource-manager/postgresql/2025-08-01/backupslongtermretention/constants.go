package backupslongtermretention

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *ExecutionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExecutionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
