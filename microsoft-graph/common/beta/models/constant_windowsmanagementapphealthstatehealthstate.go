package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagementAppHealthStateHealthState string

const (
	WindowsManagementAppHealthStateHealthStatehealthy   WindowsManagementAppHealthStateHealthState = "Healthy"
	WindowsManagementAppHealthStateHealthStateunhealthy WindowsManagementAppHealthStateHealthState = "Unhealthy"
	WindowsManagementAppHealthStateHealthStateunknown   WindowsManagementAppHealthStateHealthState = "Unknown"
)

func PossibleValuesForWindowsManagementAppHealthStateHealthState() []string {
	return []string{
		string(WindowsManagementAppHealthStateHealthStatehealthy),
		string(WindowsManagementAppHealthStateHealthStateunhealthy),
		string(WindowsManagementAppHealthStateHealthStateunknown),
	}
}

func parseWindowsManagementAppHealthStateHealthState(input string) (*WindowsManagementAppHealthStateHealthState, error) {
	vals := map[string]WindowsManagementAppHealthStateHealthState{
		"healthy":   WindowsManagementAppHealthStateHealthStatehealthy,
		"unhealthy": WindowsManagementAppHealthStateHealthStateunhealthy,
		"unknown":   WindowsManagementAppHealthStateHealthStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagementAppHealthStateHealthState(input)
	return &out, nil
}
