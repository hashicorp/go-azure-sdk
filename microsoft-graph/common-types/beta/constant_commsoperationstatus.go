package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommsOperationStatus string

const (
	CommsOperationStatus_Completed  CommsOperationStatus = "Completed"
	CommsOperationStatus_Failed     CommsOperationStatus = "Failed"
	CommsOperationStatus_NotStarted CommsOperationStatus = "NotStarted"
	CommsOperationStatus_Running    CommsOperationStatus = "Running"
)

func PossibleValuesForCommsOperationStatus() []string {
	return []string{
		string(CommsOperationStatus_Completed),
		string(CommsOperationStatus_Failed),
		string(CommsOperationStatus_NotStarted),
		string(CommsOperationStatus_Running),
	}
}

func (s *CommsOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCommsOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCommsOperationStatus(input string) (*CommsOperationStatus, error) {
	vals := map[string]CommsOperationStatus{
		"completed":  CommsOperationStatus_Completed,
		"failed":     CommsOperationStatus_Failed,
		"notstarted": CommsOperationStatus_NotStarted,
		"running":    CommsOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommsOperationStatus(input)
	return &out, nil
}
