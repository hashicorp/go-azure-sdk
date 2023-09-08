package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeploymentStateRequestedValue string

const (
	WindowsUpdatesDeploymentStateRequestedValuearchived WindowsUpdatesDeploymentStateRequestedValue = "Archived"
	WindowsUpdatesDeploymentStateRequestedValuenone     WindowsUpdatesDeploymentStateRequestedValue = "None"
	WindowsUpdatesDeploymentStateRequestedValuepaused   WindowsUpdatesDeploymentStateRequestedValue = "Paused"
)

func PossibleValuesForWindowsUpdatesDeploymentStateRequestedValue() []string {
	return []string{
		string(WindowsUpdatesDeploymentStateRequestedValuearchived),
		string(WindowsUpdatesDeploymentStateRequestedValuenone),
		string(WindowsUpdatesDeploymentStateRequestedValuepaused),
	}
}

func parseWindowsUpdatesDeploymentStateRequestedValue(input string) (*WindowsUpdatesDeploymentStateRequestedValue, error) {
	vals := map[string]WindowsUpdatesDeploymentStateRequestedValue{
		"archived": WindowsUpdatesDeploymentStateRequestedValuearchived,
		"none":     WindowsUpdatesDeploymentStateRequestedValuenone,
		"paused":   WindowsUpdatesDeploymentStateRequestedValuepaused,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesDeploymentStateRequestedValue(input)
	return &out, nil
}
