package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageRightState string

const (
	UsageRightStateactive    UsageRightState = "Active"
	UsageRightStateinactive  UsageRightState = "Inactive"
	UsageRightStatesuspended UsageRightState = "Suspended"
	UsageRightStatewarning   UsageRightState = "Warning"
)

func PossibleValuesForUsageRightState() []string {
	return []string{
		string(UsageRightStateactive),
		string(UsageRightStateinactive),
		string(UsageRightStatesuspended),
		string(UsageRightStatewarning),
	}
}

func parseUsageRightState(input string) (*UsageRightState, error) {
	vals := map[string]UsageRightState{
		"active":    UsageRightStateactive,
		"inactive":  UsageRightStateinactive,
		"suspended": UsageRightStatesuspended,
		"warning":   UsageRightStatewarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UsageRightState(input)
	return &out, nil
}
