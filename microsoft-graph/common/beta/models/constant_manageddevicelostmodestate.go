package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceLostModeState string

const (
	ManagedDeviceLostModeStatedisabled ManagedDeviceLostModeState = "Disabled"
	ManagedDeviceLostModeStateenabled  ManagedDeviceLostModeState = "Enabled"
)

func PossibleValuesForManagedDeviceLostModeState() []string {
	return []string{
		string(ManagedDeviceLostModeStatedisabled),
		string(ManagedDeviceLostModeStateenabled),
	}
}

func parseManagedDeviceLostModeState(input string) (*ManagedDeviceLostModeState, error) {
	vals := map[string]ManagedDeviceLostModeState{
		"disabled": ManagedDeviceLostModeStatedisabled,
		"enabled":  ManagedDeviceLostModeStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceLostModeState(input)
	return &out, nil
}
