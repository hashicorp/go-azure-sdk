package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceActivityState string

const (
	TeamworkDeviceActivityStatebusy        TeamworkDeviceActivityState = "Busy"
	TeamworkDeviceActivityStateidle        TeamworkDeviceActivityState = "Idle"
	TeamworkDeviceActivityStateunavailable TeamworkDeviceActivityState = "Unavailable"
	TeamworkDeviceActivityStateunknown     TeamworkDeviceActivityState = "Unknown"
)

func PossibleValuesForTeamworkDeviceActivityState() []string {
	return []string{
		string(TeamworkDeviceActivityStatebusy),
		string(TeamworkDeviceActivityStateidle),
		string(TeamworkDeviceActivityStateunavailable),
		string(TeamworkDeviceActivityStateunknown),
	}
}

func parseTeamworkDeviceActivityState(input string) (*TeamworkDeviceActivityState, error) {
	vals := map[string]TeamworkDeviceActivityState{
		"busy":        TeamworkDeviceActivityStatebusy,
		"idle":        TeamworkDeviceActivityStateidle,
		"unavailable": TeamworkDeviceActivityStateunavailable,
		"unknown":     TeamworkDeviceActivityStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkDeviceActivityState(input)
	return &out, nil
}
