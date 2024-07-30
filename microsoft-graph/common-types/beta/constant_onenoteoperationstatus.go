package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteOperationStatus string

const (
	OnenoteOperationStatus_Completed  OnenoteOperationStatus = "Completed"
	OnenoteOperationStatus_Failed     OnenoteOperationStatus = "Failed"
	OnenoteOperationStatus_NotStarted OnenoteOperationStatus = "NotStarted"
	OnenoteOperationStatus_Running    OnenoteOperationStatus = "Running"
)

func PossibleValuesForOnenoteOperationStatus() []string {
	return []string{
		string(OnenoteOperationStatus_Completed),
		string(OnenoteOperationStatus_Failed),
		string(OnenoteOperationStatus_NotStarted),
		string(OnenoteOperationStatus_Running),
	}
}

func (s *OnenoteOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnenoteOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnenoteOperationStatus(input string) (*OnenoteOperationStatus, error) {
	vals := map[string]OnenoteOperationStatus{
		"completed":  OnenoteOperationStatus_Completed,
		"failed":     OnenoteOperationStatus_Failed,
		"notstarted": OnenoteOperationStatus_NotStarted,
		"running":    OnenoteOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnenoteOperationStatus(input)
	return &out, nil
}
