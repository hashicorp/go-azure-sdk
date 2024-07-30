package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordOperationStatus string

const (
	RecordOperationStatus_Completed  RecordOperationStatus = "Completed"
	RecordOperationStatus_Failed     RecordOperationStatus = "Failed"
	RecordOperationStatus_NotStarted RecordOperationStatus = "NotStarted"
	RecordOperationStatus_Running    RecordOperationStatus = "Running"
)

func PossibleValuesForRecordOperationStatus() []string {
	return []string{
		string(RecordOperationStatus_Completed),
		string(RecordOperationStatus_Failed),
		string(RecordOperationStatus_NotStarted),
		string(RecordOperationStatus_Running),
	}
}

func (s *RecordOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecordOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecordOperationStatus(input string) (*RecordOperationStatus, error) {
	vals := map[string]RecordOperationStatus{
		"completed":  RecordOperationStatus_Completed,
		"failed":     RecordOperationStatus_Failed,
		"notstarted": RecordOperationStatus_NotStarted,
		"running":    RecordOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecordOperationStatus(input)
	return &out, nil
}
