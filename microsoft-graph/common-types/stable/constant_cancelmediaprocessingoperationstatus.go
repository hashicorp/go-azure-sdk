package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CancelMediaProcessingOperationStatus string

const (
	CancelMediaProcessingOperationStatus_Completed  CancelMediaProcessingOperationStatus = "Completed"
	CancelMediaProcessingOperationStatus_Failed     CancelMediaProcessingOperationStatus = "Failed"
	CancelMediaProcessingOperationStatus_NotStarted CancelMediaProcessingOperationStatus = "NotStarted"
	CancelMediaProcessingOperationStatus_Running    CancelMediaProcessingOperationStatus = "Running"
)

func PossibleValuesForCancelMediaProcessingOperationStatus() []string {
	return []string{
		string(CancelMediaProcessingOperationStatus_Completed),
		string(CancelMediaProcessingOperationStatus_Failed),
		string(CancelMediaProcessingOperationStatus_NotStarted),
		string(CancelMediaProcessingOperationStatus_Running),
	}
}

func (s *CancelMediaProcessingOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCancelMediaProcessingOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCancelMediaProcessingOperationStatus(input string) (*CancelMediaProcessingOperationStatus, error) {
	vals := map[string]CancelMediaProcessingOperationStatus{
		"completed":  CancelMediaProcessingOperationStatus_Completed,
		"failed":     CancelMediaProcessingOperationStatus_Failed,
		"notstarted": CancelMediaProcessingOperationStatus_NotStarted,
		"running":    CancelMediaProcessingOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CancelMediaProcessingOperationStatus(input)
	return &out, nil
}
