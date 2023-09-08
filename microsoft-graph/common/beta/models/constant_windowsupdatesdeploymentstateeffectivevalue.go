package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeploymentStateEffectiveValue string

const (
	WindowsUpdatesDeploymentStateEffectiveValuearchived  WindowsUpdatesDeploymentStateEffectiveValue = "Archived"
	WindowsUpdatesDeploymentStateEffectiveValuefaulted   WindowsUpdatesDeploymentStateEffectiveValue = "Faulted"
	WindowsUpdatesDeploymentStateEffectiveValueoffering  WindowsUpdatesDeploymentStateEffectiveValue = "Offering"
	WindowsUpdatesDeploymentStateEffectiveValuepaused    WindowsUpdatesDeploymentStateEffectiveValue = "Paused"
	WindowsUpdatesDeploymentStateEffectiveValuescheduled WindowsUpdatesDeploymentStateEffectiveValue = "Scheduled"
)

func PossibleValuesForWindowsUpdatesDeploymentStateEffectiveValue() []string {
	return []string{
		string(WindowsUpdatesDeploymentStateEffectiveValuearchived),
		string(WindowsUpdatesDeploymentStateEffectiveValuefaulted),
		string(WindowsUpdatesDeploymentStateEffectiveValueoffering),
		string(WindowsUpdatesDeploymentStateEffectiveValuepaused),
		string(WindowsUpdatesDeploymentStateEffectiveValuescheduled),
	}
}

func parseWindowsUpdatesDeploymentStateEffectiveValue(input string) (*WindowsUpdatesDeploymentStateEffectiveValue, error) {
	vals := map[string]WindowsUpdatesDeploymentStateEffectiveValue{
		"archived":  WindowsUpdatesDeploymentStateEffectiveValuearchived,
		"faulted":   WindowsUpdatesDeploymentStateEffectiveValuefaulted,
		"offering":  WindowsUpdatesDeploymentStateEffectiveValueoffering,
		"paused":    WindowsUpdatesDeploymentStateEffectiveValuepaused,
		"scheduled": WindowsUpdatesDeploymentStateEffectiveValuescheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesDeploymentStateEffectiveValue(input)
	return &out, nil
}
