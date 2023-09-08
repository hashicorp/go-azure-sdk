package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotSettingsSyncStatus string

const (
	WindowsAutopilotSettingsSyncStatuscompleted  WindowsAutopilotSettingsSyncStatus = "Completed"
	WindowsAutopilotSettingsSyncStatusfailed     WindowsAutopilotSettingsSyncStatus = "Failed"
	WindowsAutopilotSettingsSyncStatusinProgress WindowsAutopilotSettingsSyncStatus = "InProgress"
	WindowsAutopilotSettingsSyncStatusunknown    WindowsAutopilotSettingsSyncStatus = "Unknown"
)

func PossibleValuesForWindowsAutopilotSettingsSyncStatus() []string {
	return []string{
		string(WindowsAutopilotSettingsSyncStatuscompleted),
		string(WindowsAutopilotSettingsSyncStatusfailed),
		string(WindowsAutopilotSettingsSyncStatusinProgress),
		string(WindowsAutopilotSettingsSyncStatusunknown),
	}
}

func parseWindowsAutopilotSettingsSyncStatus(input string) (*WindowsAutopilotSettingsSyncStatus, error) {
	vals := map[string]WindowsAutopilotSettingsSyncStatus{
		"completed":  WindowsAutopilotSettingsSyncStatuscompleted,
		"failed":     WindowsAutopilotSettingsSyncStatusfailed,
		"inprogress": WindowsAutopilotSettingsSyncStatusinProgress,
		"unknown":    WindowsAutopilotSettingsSyncStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotSettingsSyncStatus(input)
	return &out, nil
}
