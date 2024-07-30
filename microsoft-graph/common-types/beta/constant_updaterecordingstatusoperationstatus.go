package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRecordingStatusOperationStatus string

const (
	UpdateRecordingStatusOperationStatus_Completed  UpdateRecordingStatusOperationStatus = "Completed"
	UpdateRecordingStatusOperationStatus_Failed     UpdateRecordingStatusOperationStatus = "Failed"
	UpdateRecordingStatusOperationStatus_NotStarted UpdateRecordingStatusOperationStatus = "NotStarted"
	UpdateRecordingStatusOperationStatus_Running    UpdateRecordingStatusOperationStatus = "Running"
)

func PossibleValuesForUpdateRecordingStatusOperationStatus() []string {
	return []string{
		string(UpdateRecordingStatusOperationStatus_Completed),
		string(UpdateRecordingStatusOperationStatus_Failed),
		string(UpdateRecordingStatusOperationStatus_NotStarted),
		string(UpdateRecordingStatusOperationStatus_Running),
	}
}

func (s *UpdateRecordingStatusOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdateRecordingStatusOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdateRecordingStatusOperationStatus(input string) (*UpdateRecordingStatusOperationStatus, error) {
	vals := map[string]UpdateRecordingStatusOperationStatus{
		"completed":  UpdateRecordingStatusOperationStatus_Completed,
		"failed":     UpdateRecordingStatusOperationStatus_Failed,
		"notstarted": UpdateRecordingStatusOperationStatus_NotStarted,
		"running":    UpdateRecordingStatusOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateRecordingStatusOperationStatus(input)
	return &out, nil
}
