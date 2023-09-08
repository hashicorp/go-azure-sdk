package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationNotificationTargettedUserType string

const (
	SimulationNotificationTargettedUserTypeallUsers    SimulationNotificationTargettedUserType = "AllUsers"
	SimulationNotificationTargettedUserTypeclicked     SimulationNotificationTargettedUserType = "Clicked"
	SimulationNotificationTargettedUserTypecompromised SimulationNotificationTargettedUserType = "Compromised"
	SimulationNotificationTargettedUserTypeunknown     SimulationNotificationTargettedUserType = "Unknown"
)

func PossibleValuesForSimulationNotificationTargettedUserType() []string {
	return []string{
		string(SimulationNotificationTargettedUserTypeallUsers),
		string(SimulationNotificationTargettedUserTypeclicked),
		string(SimulationNotificationTargettedUserTypecompromised),
		string(SimulationNotificationTargettedUserTypeunknown),
	}
}

func parseSimulationNotificationTargettedUserType(input string) (*SimulationNotificationTargettedUserType, error) {
	vals := map[string]SimulationNotificationTargettedUserType{
		"allusers":    SimulationNotificationTargettedUserTypeallUsers,
		"clicked":     SimulationNotificationTargettedUserTypeclicked,
		"compromised": SimulationNotificationTargettedUserTypecompromised,
		"unknown":     SimulationNotificationTargettedUserTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationNotificationTargettedUserType(input)
	return &out, nil
}
