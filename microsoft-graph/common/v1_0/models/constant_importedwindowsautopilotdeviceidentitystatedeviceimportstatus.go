package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus string

const (
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatuscomplete ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "Complete"
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatuserror    ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "Error"
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatuspartial  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "Partial"
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatuspending  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "Pending"
	ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusunknown  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus = "Unknown"
)

func PossibleValuesForImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus() []string {
	return []string{
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatuscomplete),
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatuserror),
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatuspartial),
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatuspending),
		string(ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusunknown),
	}
}

func parseImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus(input string) (*ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus, error) {
	vals := map[string]ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus{
		"complete": ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatuscomplete,
		"error":    ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatuserror,
		"partial":  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatuspartial,
		"pending":  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatuspending,
		"unknown":  ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus(input)
	return &out, nil
}
