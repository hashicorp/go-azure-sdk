package importedwindowsautopilotdeviceidentity

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus string

const (
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusComplete ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "complete"
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusError    ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "error"
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusPartial  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "partial"
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusPending  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "pending"
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusUnknown  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "unknown"
)

func PossibleValuesForImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus() []string {
	return []string{
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusComplete),
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusError),
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusPartial),
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusPending),
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusUnknown),
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
		"complete": ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusComplete,
		"error":    ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusError,
		"partial":  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusPartial,
		"pending":  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusPending,
		"unknown":  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus(input)
	return &out, nil
}
