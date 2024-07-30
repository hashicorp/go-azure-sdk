package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotSettingsSyncStatus string

const (
	WindowsAutopilotSettingsSyncStatus_Completed  WindowsAutopilotSettingsSyncStatus = "completed"
	WindowsAutopilotSettingsSyncStatus_Failed     WindowsAutopilotSettingsSyncStatus = "failed"
	WindowsAutopilotSettingsSyncStatus_InProgress WindowsAutopilotSettingsSyncStatus = "inProgress"
	WindowsAutopilotSettingsSyncStatus_Unknown    WindowsAutopilotSettingsSyncStatus = "unknown"
)

func PossibleValuesForWindowsAutopilotSettingsSyncStatus() []string {
	return []string{
		string(WindowsAutopilotSettingsSyncStatus_Completed),
		string(WindowsAutopilotSettingsSyncStatus_Failed),
		string(WindowsAutopilotSettingsSyncStatus_InProgress),
		string(WindowsAutopilotSettingsSyncStatus_Unknown),
	}
}

func (s *WindowsAutopilotSettingsSyncStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotSettingsSyncStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotSettingsSyncStatus(input string) (*WindowsAutopilotSettingsSyncStatus, error) {
	vals := map[string]WindowsAutopilotSettingsSyncStatus{
		"completed":  WindowsAutopilotSettingsSyncStatus_Completed,
		"failed":     WindowsAutopilotSettingsSyncStatus_Failed,
		"inprogress": WindowsAutopilotSettingsSyncStatus_InProgress,
		"unknown":    WindowsAutopilotSettingsSyncStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotSettingsSyncStatus(input)
	return &out, nil
}
