package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus string

const (
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Complete ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "complete"
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Error    ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "error"
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Partial  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "partial"
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Pending  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "pending"
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Unknown  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "unknown"
)

func PossibleValuesForImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus() []string {
	return []string{
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Complete),
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Error),
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Partial),
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Pending),
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Unknown),
	}
}

func (s *ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus(input string) (*ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus, error) {
	vals := map[string]ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus{
		"complete": ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Complete,
		"error":    ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Error,
		"partial":  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Partial,
		"pending":  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Pending,
		"unknown":  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus(input)
	return &out, nil
}
