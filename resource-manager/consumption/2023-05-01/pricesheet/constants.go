package pricesheet

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatusType string

const (
	OperationStatusTypeCompleted OperationStatusType = "Completed"
	OperationStatusTypeFailed    OperationStatusType = "Failed"
	OperationStatusTypeRunning   OperationStatusType = "Running"
)

func PossibleValuesForOperationStatusType() []string {
	return []string{
		string(OperationStatusTypeCompleted),
		string(OperationStatusTypeFailed),
		string(OperationStatusTypeRunning),
	}
}

func (s *OperationStatusType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationStatusType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationStatusType(input string) (*OperationStatusType, error) {
	vals := map[string]OperationStatusType{
		"completed": OperationStatusTypeCompleted,
		"failed":    OperationStatusTypeFailed,
		"running":   OperationStatusTypeRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationStatusType(input)
	return &out, nil
}
