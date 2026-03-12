package migratemysqlstatusoperationgroup

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatus string

const (
	OperationStatusCreated    OperationStatus = "Created"
	OperationStatusFailed     OperationStatus = "Failed"
	OperationStatusInProgress OperationStatus = "InProgress"
	OperationStatusSucceeded  OperationStatus = "Succeeded"
	OperationStatusTimedOut   OperationStatus = "TimedOut"
)

func PossibleValuesForOperationStatus() []string {
	return []string{
		string(OperationStatusCreated),
		string(OperationStatusFailed),
		string(OperationStatusInProgress),
		string(OperationStatusSucceeded),
		string(OperationStatusTimedOut),
	}
}

func (s *OperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationStatus(input string) (*OperationStatus, error) {
	vals := map[string]OperationStatus{
		"created":    OperationStatusCreated,
		"failed":     OperationStatusFailed,
		"inprogress": OperationStatusInProgress,
		"succeeded":  OperationStatusSucceeded,
		"timedout":   OperationStatusTimedOut,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationStatus(input)
	return &out, nil
}
