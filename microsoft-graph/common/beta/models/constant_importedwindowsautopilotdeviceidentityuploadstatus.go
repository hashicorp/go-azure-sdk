package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedWindowsAutopilotDeviceIdentityUploadStatus string

const (
	ImportedWindowsAutopilotDeviceIdentityUploadStatuscomplete ImportedWindowsAutopilotDeviceIdentityUploadStatus = "Complete"
	ImportedWindowsAutopilotDeviceIdentityUploadStatuserror    ImportedWindowsAutopilotDeviceIdentityUploadStatus = "Error"
	ImportedWindowsAutopilotDeviceIdentityUploadStatusnoUpload ImportedWindowsAutopilotDeviceIdentityUploadStatus = "NoUpload"
	ImportedWindowsAutopilotDeviceIdentityUploadStatuspending  ImportedWindowsAutopilotDeviceIdentityUploadStatus = "Pending"
)

func PossibleValuesForImportedWindowsAutopilotDeviceIdentityUploadStatus() []string {
	return []string{
		string(ImportedWindowsAutopilotDeviceIdentityUploadStatuscomplete),
		string(ImportedWindowsAutopilotDeviceIdentityUploadStatuserror),
		string(ImportedWindowsAutopilotDeviceIdentityUploadStatusnoUpload),
		string(ImportedWindowsAutopilotDeviceIdentityUploadStatuspending),
	}
}

func parseImportedWindowsAutopilotDeviceIdentityUploadStatus(input string) (*ImportedWindowsAutopilotDeviceIdentityUploadStatus, error) {
	vals := map[string]ImportedWindowsAutopilotDeviceIdentityUploadStatus{
		"complete": ImportedWindowsAutopilotDeviceIdentityUploadStatuscomplete,
		"error":    ImportedWindowsAutopilotDeviceIdentityUploadStatuserror,
		"noupload": ImportedWindowsAutopilotDeviceIdentityUploadStatusnoUpload,
		"pending":  ImportedWindowsAutopilotDeviceIdentityUploadStatuspending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedWindowsAutopilotDeviceIdentityUploadStatus(input)
	return &out, nil
}
